package piscine

// BasicAtoi2 : ba2
func BasicAtoi2(s string) int {
	res := 0
	if isNumeric(s) {
		for _, el := range s {
			res = res*10 + int(el-48)
		}
		return res
	}
	return 0

}
func isNumeric(s string) bool {
	for _, el := range s {
		if el == '+' || el == '-' {
			continue
		}
		if el < 48 || el > 57 {
			return false
		}
	}
	return true

}
