package piscine

// IsNumeric : isnumeric

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	if str[0] == '+' || str[0] == '-' {
		str = str[1:]
	}
	for _, letter := range str {
		if letter < '0' || letter > '9' {
			return false
		}
	}
	return true
}
