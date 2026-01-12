func maxSubArray(nums []int) int {
    return divideAndConquer(nums, 0, len(nums)-1)
}

func divideAndConquer(nums []int, left, right int) int {
  
    if left == right {
        return nums[left]
    }

    mid := (left + right) / 2

    // 1. Tìm Max bên trái (đệ quy)
    leftMax := divideAndConquer(nums, left, mid)
    
   
    rightMax := divideAndConquer(nums, mid+1, right)
    
   
    crossMax := maxCrossingSum(nums, left, mid, right)

   
    return max(leftMax, max(rightMax, crossMax))
}

func maxCrossingSum(nums []int, left, mid, right int) int {
   
    sum := 0
    leftPart := -100000 
    for i := mid; i >= left; i-- {
        sum += nums[i]
        if sum > leftPart {
            leftPart = sum
        }
    }

    
    sum = 0
    rightPart := -100000
    for i := mid + 1; i <= right; i++ {
        sum += nums[i]
        if sum > rightPart {
            rightPart = sum
        }
    }

    return leftPart + rightPart
}

func max(a, b int) int {
    if a > b { return a }
    return b
}