func isPalindrome(x int) bool {
	if x <0 || (x%10 ==0 &&  x!=0) {
        return false
    }
    orgin := x
    reversed := 0
   for x > 0 {
       y := x%10
       reversed = y + reversed*10
       x=x/10
   }
   return reversed == orgin
    
}