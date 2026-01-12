func containsDuplicate(nums []int) bool {
    l := len(nums)
    for i := 1; i < l; i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    
    sort.Ints(nums)
    for i := 1; i < l; i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }

    return false
}