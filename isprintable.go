package piscine

func IsPrintable(s string) bool {
	s1bool := true
	s2Rune := []rune(s)
	for _, key := range s2Rune {
		if key >= 0 && key <= 31 {
			s1bool = false
		}
	}
	return s1bool
}
