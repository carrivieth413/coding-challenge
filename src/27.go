func calculateGCD(a, b int) int {
    if b == 0 {
        return a
    }
    return calculateGCD(b, a%b)
}
