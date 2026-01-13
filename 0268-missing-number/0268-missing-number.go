func missingNumber(nums []int) int {
    n := len(nums)
    visited := make([]bool, n+1)

    for _,v := range nums{
        visited[v] = true
    }

    for i:=0 ; i<= n;i++{
        if visited[i] == false{
            return i
        }
    }
    return -1
}