package piscine

func IsPrintable(s string) bool {
	sbool := true
	sRune := []rune(s)
	for _, key := range sRune {
		if key >= 0 && key >= 31 {
			sbool = false
		}
	}
	return stbool
}
