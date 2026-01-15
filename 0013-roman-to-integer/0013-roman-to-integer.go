func romanToInt(s string) int {
  m := map[byte]int{
    'I' : 1,
    'V' : 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
  }

  value := 0
  n := len(s)
  for i:=0 ; i< n; i++{
    currentVal := m[s[i]]
    if i < n -1 && currentVal < m[s[i+1]] {
        value -= currentVal
    }else {
        value += currentVal
    }  
  }
  return value
}