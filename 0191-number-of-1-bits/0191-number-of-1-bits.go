func hammingWeight(n int) int {
    count := 0
    for n > 0 {
        n &= (n - 1) // Xóa phắt bit 1 cuối cùng
        count++      // Mỗi lần xóa được thì đếm 1
    }
    return count
}