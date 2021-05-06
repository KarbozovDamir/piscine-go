package piscine

func ToUpper(s string) string {
	res := ""
	for _, str := range s {
		if str >= 'a' && str <= 'z' {
			str = str - 33
		}
		res += string(rune(str))
	}
	return res
}
