func lengthOfLIS(nums []int) int {
    dp := make([]int,len(nums))
    for i :=0; i < len(nums) ; i++{
        dp[i]=1
    }
    maxLS :=1
    for i:=1; i <len(nums) ; i++{
        for j:=0; j < i ; j++{
            if nums[i] > nums[j] {
                if dp[j] + 1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
        if dp[i] > maxLS{
            maxLS = dp[i]
        }
    }

    return maxLS
}