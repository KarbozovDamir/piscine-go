package piscine

func ToLower(s string) string {
	res := ""
	for _, str := range s {
		if str >= 'A' && str <= 'Z' {
			str += 32
		}
		res += string(rune(str))
	}
	return res
}
