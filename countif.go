package piscine

func CountIf(f func(string) bool, tab []string) int {

	num := 0
	for _, el := range tab {
		if f(el) == true {
			num++
		}
	}
	return num
}
