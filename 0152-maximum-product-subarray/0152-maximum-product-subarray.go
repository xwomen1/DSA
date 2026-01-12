func maxProduct(nums []int) int {
    maxProd := nums[0]
    curentProd := nums[0]
    minProd := nums[0]

    for i:=1; i<len(nums);i++{
      v := nums[i]
      if v < 0 {
        curentProd, minProd = minProd, curentProd
      }

      curentProd = max(v, curentProd *v)
      minProd = min(v, minProd*v)

      if curentProd > maxProd {
        maxProd = curentProd
      }
        
    }
    
    return maxProd
}