package piscine

func NRune(s string, n int) rune {
	res := []rune(s)

	for index, value := range res {
		if index == n-1 {
			return value
		}
	}
	return 0
}
