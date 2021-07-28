package piscine

// BasicAtoi : ba
func BasicAtoi(s string) int {
	res := 0
	for _, el := range s {
		res = res*10 + int(el-48)
	}
	return res
}
