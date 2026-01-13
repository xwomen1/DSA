func maxArea(height []int) int {
   n := len(height)
   left := 0
   right := n-1
   maxW := 0
 
   for left < right {
    width := right - left
    hight := min(height[right],height[left])
    currentW := width*hight
    if currentW > maxW {
        maxW = currentW
    }
    if height[left] >height[right]{
        right--
    }else{left++}
    
   } 
   return maxW
}