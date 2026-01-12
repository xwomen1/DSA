func minTimeToVisitAllPoints(points [][]int) int {
   toT := 0
   for i := 0; i< len(points) -1; i++{
    p1 := points[i]
    p2 := points[i+1]

    diffX := abs(p2[0] - p1[0])
    diffY := abs(p2[1] - p1[1])

    toT += max(diffX, diffY)
   } 
   return toT
}

func abs (x int) int{
    if x <0 {
        return -x
    }
    return x
}

func max(a,b int ) int{
    if a > b {
        return a
    }
    return b
}