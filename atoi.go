package piscine

// Atoi : atoi
func Atoi(s string) int {
	res := BasicAtoi2(s)
	return res
}

// func BasicAtoi2(s string) int {
// 	if IsNumeric(s) { // если нашли число
// 		res := 0
// 		neg := false
// 		if s[0] == '+' || s[0] == '-' {
// 			if s[0] == '-' {
// 				neg = true
// 			}
// 			s = s[1:]
// 		}
// 		for _, el := range s { // собираем число в переменной - res
// 			res = res*10 + int(el-48) // разбор числа по букваим(числам, индексам)
// 		}
// 		if neg == true {
// 			res = -res
// 		}
// 		return res
// 	}
// 	return 0 // если не число тогда возвращаем ноль
// }

// func IsNumeric(str string) bool {
// 	if str == "" {
// 		return false
// 	}
// 	if str[0] == '+' || str[0] == '-' {
// 		str = str[1:]
// 	}
// 	for _, letter := range str {
// 		if letter < '0' || letter > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }
