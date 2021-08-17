package piscine

func Atoi(s string) int {
	res := 0
	neg := false
	if s[0] == '+' {
		s = s[1:]
	}
	if s[0] == '-' {
		s = s[1:]
		neg = true
	}

	for _, el := range s {
		if el >= '0' && el <= '9' {
			res = res*10 + int(el-48)
		} else {
			return 0
		}
	}
	if neg {
		res *= -1
	}
	if !neg && res < 0 {
		return 0
	}
	if neg && res > 0 {
		return 0

	}
	return res
}
