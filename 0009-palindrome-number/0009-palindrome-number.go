func isPalindrome(x int) bool {
	if x <0 || (x%10 ==0 &&  x!=0) {
        return false
    }
    reversed := 0
   for x > reversed {
       y := x%10
       reversed = y + reversed*10
       x=x/10
   }
   return reversed == x || reversed/10 == x
    
}