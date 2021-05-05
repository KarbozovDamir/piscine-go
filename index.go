package piscine

func Index(s string, toFind string) int {
	s1 := []rune(s)
	sF := []rune(toFind)
	a := 0
	b := 0
	for range s1 {
		a++
	}
	for range sF {
		b++
	}
	for i := 0; i < a-b; i++ {
		if toFind == s[i:i+b] {
			return (i)
		}
	}
	return -1
}
