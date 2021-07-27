package piscine

// BasicAtoi : ba
func BasicAtoi(s string) int {
	res := 0
	for _, el := range s {
		if string(el) != "-" && string(el) != "+" {
			res = res*10 + (int(el) - 48)
		}
	}
	return res
}
