package util

// 截断字符串
func StringSubstr(s string, pos, length int) string {
    runes := []rune(s)
    l := pos + length
    if l > len(runes) {
        l = len(runes)
    }
    return string(runes[pos:l])
}