package piscine

func SplitWhiteSpaces(str string) []string {
	count := 0
	strFound := false

	for _, char := range str {
		if char == ' ' || char == '\n' || char == '\t' {
			strFound = false
		} else if !strFound {
			strFound = true
			count++
		}
	}

	result := make([]string, count)

	fillIndex := -1
	strFound = false
	for _, char := range str {
		if char == ' ' || char == '\n' || char == '\t' {
			strFound = false
		} else {
			if !strFound {
				fillIndex++
			}
			strFound = true
			result[fillIndex] += string(char)
		}
	}
	return result
}
