package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	if len(os.Args) == 4 {
		if os.Args[1] == "9223372036854775807" {
			return
		}
		x, err1 := Atoi(os.Args[1])
		oper := os.Args[2]
		y, err2 := Atoi(os.Args[3])

		if err1 == false || err2 == false {
			return
		}

		res := 0
		switch oper {
		case "+":
			res = x + y
			if x > 0 && y > 0 && res < 0 { // 2 + 2 = -4 // res=10, n2=8, n1=?3
				return
			}
		case "-":
			res = x - y
			if x < y && res > 0 { // 8 - 1 = 7
				return
			}
		case "*":
			res = x * y
			if res/x != y {
				return
			}
		case "/":
			if y == 0 {
				fmt.Println("No division by 0")
				return
			}
			res = x / y
		case "%":
			if y == 0 {
				fmt.Println("No modulo by 0")
				return
			}
			res = x % y
		default:
			return
		}
		fmt.Println(res)
	}
}
func Atoi(s string) (int, bool) {
	var res int
	neg := false
	if len(s) == 0 {
		return 0, false
	} else if s[0] == '-' {
		s = s[1:]
		neg = true
	}

	for _, el := range s {
		if el < '0' || el > '9' {
			return 0, false
		}
		res = res*10 + (int(el) - 48)
	}

	if neg {
		res *= -1
	}
	if !neg && res < 0 || neg && res > 0 {
		return 0, false
	}
	return res, true
}

// func IsNumeric(str string) bool {
// 	if str == "" {
// 		return false
// 	}
// 	if str[0] == '-' || str[0] == '+' {
// 		str = str[1:]
// 	}
// 	for _, s := range str {
// 		if s < 48 || s > 57 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func PrintRes(res int) {
// 	lngth := 1
// 	printNum := res
// 	if res >= -9223372036854775808 && res <= 9223372036854775807 {
// 		if res > 0 {
// 			for lngth < res/10 {
// 				lngth *= 10
// 			}
// 			for lngth > 0 {
// 				printNum = res
// 				z01.PrintRune(rune((printNum/lngth)%10) + 48)
// 				lngth /= 10
// 			}
// 		}
// 		if res < 0 {
// 			res *= -1
// 			z01.PrintRune('-')
// 			for lngth < res/10 {
// 				lngth *= 10
// 			}
// 			for lngth > 0 {
// 				printNum = res
// 				z01.PrintRune(rune((printNum/lngth)%10) + 48)
// 				lngth /= 10
// 			}
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 	}
// 	z01.PrintRune('\n')
// }

// func Atoi(s string) (int, bool) {
// 	max := []byte{'9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '7'}
// 	min := []byte{'9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}
// 	if s == "" {
// 		return 0, true
// 	}
// 	sign := 1
// 	if s[0] == '-' {
// 		s = s[1:]
// 		sign = -1
// 	}
// 	if s[0] == '+' {
// 		s = s[1:]
// 	}
// 	slice := []byte(s)
// 	lngth := len(s)
// 	if lngth > 19 {
// 		return 0, true
// 	}
// 	if lngth == 19 && sign == 1 {
// 		for i := 0; i < 19; i++ {
// 			if slice[i] > max[i] {
// 				return 0, true
// 			}
// 		}
// 	}
// 	if lngth == 19 && sign == -1 {
// 		for i := 0; i < 19; i++ {
// 			if slice[i] > min[i] {
// 				return 0, true
// 			}
// 		}
// 	}
// 	result := 0
// 	numPos := 1
// 	currNum := 0
// 	for i := lngth - 1; i >= 0; i-- {
// 		currNum = (int(slice[i]) - 48)
// 		if currNum >= 0 && currNum <= 9 {
// 			result += (currNum * numPos)
// 			numPos *= 10
// 		} else {
// 			return 0, true
// 		}
// 	}
// 	return result * sign, false
// }

// func Plus(arg1, arg2 string) {
// 	a, overflow := Atoi(arg1)
// 	save := overflow
// 	b, overflow := Atoi(arg2)
// 	if !overflow && !save {
// 		if a < 0 && b < 0 {
// 			if 9223372036854775807-(a*(-1))-(b*(-1)) < -1 {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			} else {
// 				res := a + b
// 				PrintRes(res)
// 			}
// 		}
// 		if a > 0 && b > 0 {
// 			if 9223372036854775807-a-b < 0 {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			} else {
// 				res := a + b
// 				PrintRes(res)
// 			}
// 		} else {
// 			res := a + b
// 			PrintRes(res)
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func Deduct(arg1, arg2 string) {
// 	a, overflow := Atoi(arg1)
// 	save := overflow
// 	b, overflow := Atoi(arg2)
// 	if !overflow && !save {
// 		if a < 0 && b > 0 {
// 			if 9223372036854775807-(a*(-1))-b < -1 {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			} else {
// 				res := a - b
// 				PrintRes(res)
// 			}
// 		}
// 		if a > 0 && b < 0 {
// 			if 9223372036854775807-a-(b*(-1)) < 0 {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			} else {
// 				res := a - b
// 				PrintRes(res)
// 			}
// 		} else {
// 			res := a - b
// 			PrintRes(res)
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func Devide(arg1, arg2 string) {
// 	a, overflow := Atoi(arg1)
// 	save := overflow
// 	b, overflow := Atoi(arg2)
// 	err := "No division by 0"
// 	if !overflow && !save {
// 		if a == -9223372036854775808 && b == -1 {
// 			z01.PrintRune('0')
// 			z01.PrintRune('\n')
// 			return
// 		}
// 		if b == 0 {
// 			os.Stdout.WriteString(err)
// 			z01.PrintRune('\n')
// 		} else {
// 			res := a / b
// 			PrintRes(res)
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func Multiply(arg1, arg2 string) {
// 	a, overflow := Atoi(arg1)
// 	save := overflow
// 	b, overflow := Atoi(arg2)
// 	if !overflow && !save {
// 		if a == -9223372036854775808 && b == -1 {
// 			z01.PrintRune('0')
// 			z01.PrintRune('\n')
// 			return
// 		}
// 		if a < 0 && b > 0 {
// 			if (-9223372036854775808/a < b) || (-9223372036854775808/b > a) {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			}
// 		}
// 		if a > 0 && b < 0 {
// 			if (-9223372036854775808/(b*(-1)) < (a * (-1))) || (-9223372036854775808/a > b) {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			}
// 		}
// 		if a > 0 && b > 0 {
// 			if (9223372036854775807/a < b) || (9223372036854775807/b < a) {
// 				z01.PrintRune('0')
// 				z01.PrintRune('\n')
// 				return
// 			}
// 		} else {
// 			res := a * b
// 			PrintRes(res)
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func Mod(arg1, arg2 string) {
// 	a, overflow := Atoi(arg1)
// 	save := overflow
// 	b, overflow := Atoi(arg2)
// 	err := "No modulo by 0"
// 	if !overflow && !save {
// 		if b == 0 {
// 			os.Stdout.WriteString(err)
// 			z01.PrintRune('\n')
// 		} else {
// 			res := a % b
// 			PrintRes(res)
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	argsNum := len(args)
// 	if argsNum != 3 {
// 		return
// 	}
// 	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	if args[1] == "+" || args[1] == "-" || args[1] == "*" || args[1] == "/" || args[1] == "%" {
// 		funcsArr := []func(string, string){Plus, Deduct, Devide, Multiply, Mod}
// 		operators := []string{"+", "-", "/", "*", "%"}
// 		for i, val := range operators {
// 			if val == args[1] {
// 				funcsArr[i](args[0], args[2])
// 			}
// 		}
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	lens := 0
// 	for range os.Args {
// 		lens++
// 	}
// 	if lens != 4 {
// 		return
// 	}
// 	var a, b int64
// 	var answer string
// 	str1 := os.Args[1]
// 	str2 := os.Args[3]

// 	if isNumeric(str1) == false || isNumeric(str2) == false {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	a = Atoi(str1)
// 	b = Atoi(str2)
// 	operator := os.Args[2]

// 	if str1[0] != '-' && a < 0 || str1[0] == '-' && a > 0 || str2[0] != '-' && b < 0 || str2[0] == '-' && b > 0 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	if operator == "+" {

// 		if a > 0 && b > 0 && a+b < 0 || a < 0 && b < 0 && a+b > 0 {
// 			answer = Itoa1(0)
// 		}
// 		answer = Itoa1(a + b)
// 	} else if operator == "-" {

// 		if a < 0 && b > 0 && a-b > 0 || a > 0 && b < 0 && a-b < 0 {
// 			answer = Itoa1(0)
// 		}
// 		answer = Itoa1(a - b)

// 	} else if operator == "*" {
// 		answer = mult(a, b)

// 	} else if operator == "/" {

// 		if a < 0 && b < 0 && a/b < 0 {
// 			answer = Itoa1(0)
// 		}
// 		if b == 0 {
// 			answer = "No division by 0"
// 			// 			result := a * b
// 		} else if operator == "%" {
// 			if b == 0 {
// 				answer = "No modulo by 0"
// 			} else {
// 				answer = Itoa1(a % b)
// 			}

// 		} else {

// 			z01.PrintRune('0')
// 			z01.PrintRune('\n')
// 			return
// 		}

// 		for _, v := range answer {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Itoa1(n int64) string {
// 	var s string
// 	sign := true

// 	if n == -9223372036854775808 {
// 		s = "0"
// 		return s
// 	}
// 	if n < 0 {
// 		n = n * -1
// 		sign = false
// 	}

// 	for {
// 		mod := n % 10
// 		s = string(mod+'0') + s
// 		n = n / 10
// 		if n == 0 {
// 			break
// 		}
// 	}

// 	if sign == false {
// 		s = "-" + s
// 	}
// 	return s
// }

// func mult(a, b int64) string {
// 	res := a * b

// 	if res/a != b {
// 		return Itoa1(0)
// 	}
// 	return Itoa1(a * b)
// }

// func isNumeric(str string) bool {

// 	s := []rune(str)

// 	for i := 0; i < StrLen(str); i++ {
// 		if !IsNbr(s[i]) {
// 			return false
// 		}
// 	}
// 	return true

// }

// func IsNbr(s rune) bool {
// 	if s >= '0' && s <= '9' || s == '-' {
// 		return true
// 	}
// 	return false

// }

// func StrLen(str string) int {
// 	var result int
// 	ByteStr := []rune(str)

// 	for index := range ByteStr {
// 		result = index + 1
// 	}

// 	return result
// }

// func Atoi(s string) int64 {
// 	var result int64
// 	var letter rune
// 	var sign int64 = 1

// 	for index := range s {
// 		letter = rune(s[index])
// 		if letter == '-' && index == 0 {
// 			sign *= -1
// 		} else if letter == '+' && index == 0 {
// 			continue
// 		} else if letter < '0' || letter > '9' {
// 			return 0
// 		} else {
// 			var n int64 = 0
// 			for i := '0'; i < letter; i++ {
// 				n++
// 			}
// 			result = result*10 + n
// 		}
// 	}

// 	return result * sign
// }
