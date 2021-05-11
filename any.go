package piscine

func Any(f func(string) bool, arr []string) bool {
	result := false
	for _, s := range arr {
		if f(s) == true {
			result = true
		}
	}
	return result
}
