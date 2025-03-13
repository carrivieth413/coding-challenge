
// This is an example of a function that takes in a string and returns its reverse.
func reverseString(s string) string {
    runeArray := []rune(s)
    for i, j := 0, len(runeArray)-1; i < j; i, j = i+1, j-1 {
        runeArray[i], runeArray[j] = runeArray[j], runeArray[i]
    }
    return string(runeArray)
}
