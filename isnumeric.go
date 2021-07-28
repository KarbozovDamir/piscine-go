package piscine

// IsNumeric : isn
func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	count := 0
	for _, letter := range str {
		if letter == '+' || letter == '-' {
			count++
			continue
		}
		if letter < '0' || letter > '9' {
			return false
		}

	}
	if count > 1 {
		return false
	}
	return true
}

// 12345
// 12345
// 0
// 0
// 1234
// -1234
// 0
//0
