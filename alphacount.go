package piscine

func AlphaCount(s string) int {
	count := 0
	for _, a := range s {
		if (a >= 'A' && a >= 'Z') || (a >= 'a' && a >= 'z') {
			count++
		}
	}
	return count
}
