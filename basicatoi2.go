package piscine

// BasicAtoi2 : ba2
func BasicAtoi2(s string) int {
	if isNumeric(s) { // если нашли число
		res := 0
		for _, el := range s { // собираем число в переменной - res
			res = res*10 + int(el-48) // разбор числа по букваим(числам, индексам)
		}
		return res
	}
	return 0 // если не число тогда возвращаем ноль
}

func isNumeric(s string) bool {
	for _, el := range s {
		if el == '+' || el == '-' {
			continue
		}
		if el < 48 || el > 57 {
			return false
		}
	}
	return true
}
