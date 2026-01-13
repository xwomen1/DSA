func countBits(n int) []int {
   result := make([]int, n+1)
   count :=0
   for i :=0 ; i <= n ;i++{
    num := i
    count =0
   for num>0 {
    if num &1 ==1{count ++}
   
  num =num>>1 
   }
   result[i] =  count
   
   }
   return result
}