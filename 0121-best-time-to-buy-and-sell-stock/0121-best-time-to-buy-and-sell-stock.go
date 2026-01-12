func maxProfit(prices []int) int {
    if len(prices) == 0{
        return 0
    }

    minPrices := prices[0]
    maxProf := 0
    for i:=1; i< len(prices); i++{
        if prices[i] < minPrices {
            minPrices = prices[i]
        }else {
            cur := prices[i] - minPrices
            if cur > maxProf {
                maxProf = cur
            }
        }
    }
    return maxProf

}

