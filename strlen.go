package piscine

// func StrLen(s string) int {
// 	return len([]rune(s))
// }
func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
