package piscine

// BasicAtoi2 : ba2
func BasicAtoi2(s string) int {
	if IsNumeric(s) { // если нашли число
		res := 0
		neg := false
		if s[0] == '+' || s[0] == '-' {
			if s[0] == '-' {
				neg = true
			}
			s = s[1:]
		}
		for _, el := range s { // собираем число в переменной - res
			res = res*10 + int(el-48) // разбор числа по букваим(числам, индексам)
		}
		if neg == true {
			res = -res
		}
		return res
	}
	return 0 // если не число тогда возвращаем ноль
}
