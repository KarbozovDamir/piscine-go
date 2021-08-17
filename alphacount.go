package piscine

func AlphaCount(s string) int {
	count := 0
	for _, j := range s {
		if (j >= 'A' && j <= 'Z') || (j >= 'a' && j <= 'z') {
			count++
		}
	}
	return count
}
