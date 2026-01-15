func romanToInt(s string) int {
  var values [128]int
    values['I'] = 1
    values['V'] = 5
    values['X'] = 10
    values['L'] = 50
    values['C'] = 100
    values['D'] = 500
    values['M'] = 1000

    sum := 0
    n := len(s)
    for i := 0; i < n; i++ {
        val := values[s[i]]
        if i+1 < n && val < values[s[i+1]] {
            sum -= val
        } else {
            sum += val
        }
    }
    return sum
}