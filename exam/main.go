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

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	if len(os.Args) == 2 {
// 		n, err := strconv.Atoi(os.Args[1])
// 		if err != nil {
// 			return
// 		}
// 		if n == 1 || n == 2 {
// 			out("true")
// 			return
// 		}
// 		for n > 1 {
// 			if n%2 == 1 {
// 				out("false")
// 				return
// 			}
// 			n /= 2
// 		}
// 		out("ture")
// 	}
// }
// func out(s string) {
// 	for _, el := range s {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

//switchcase
// func main() {
// 	args := os.Args[1:]
// 	res := ""
// 	for _, el := range args[0] {
// 		res = string(el) + res
// 	}
// 	for _, el := range res {
// 		z01.PrintRune(el)

// 	}
// 	z01.PrintRune(10)
// }

//displaylastparam
// func main() {
// 	args := os.Args[1:]
// 	for _, el := range args[1] {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

// 	for i := len(args) - 1; i > 0; i-- {

// 		z01.PrintRune(rune(i))
// 	}
// 	z01.PrintRune(10)
// }

//displayfirstparamfunc main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	if len(os.Args) == 2 {
// 		n, err := strconv.Atoi(os.Args[1])
// 		if err != nil {
// 			return
// 		}
// 		if n == 1 || n == 2 {
// 			out("true")
// 			return
// 		}
// 		for n > 1 {
// 			if n%2 == 1 {
// 				out("false")
// 				return
// 			}
// 			n /= 2
// 		}
// 		out("ture")
// 	}
// }
// func out(s string) {
// 	for _, el := range s {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }
// 	for _, el := range args[0] {
// 		z01.PrintRune(el)
// 	}
// }

//hello
// func main() {
// 	s := "Hello World!"
// 	for _, el := range s {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

//**********************SwapBits
// import "fmt"

// func main() {
// 	x := byte(1)
// 	fmt.Println(SwapBits(x))
// }
// func SwapBits(octet byte) byte {
// 	return (octet>>4 | octet<<4)
// }

// **********************************Swap
// "github.com/01-edu/z01"

// func main() {
// 	a := 0
// 	b := 1
// 	Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func Swap(a *int, b *int) {
// 	c := *a
// 	*a = *b
// 	*b = c
// }

// func main() {
// 	arg := os.Args[1:]

// 	if len(arg) != 3 {
// 		return
// 	}
// 	x, err1 := Atoi(arg[0])
// 	y, err2 := Atoi(arg[2])

// 	if err1 == true || err2 == true {
// 		return
// 	}
// 	if arg[1] != "+" && arg[1] != "-" && arg[1] != "*" && arg[1] != "/" && arg[1] != "%" {
// 		return
// 	}
// 	if x == 0 && len(arg[0]) > 1 {
// 		fmt.Println(0)
// 		return
// 	}
// 	if y == 0 && len(arg[2]) > 1 {
// 		fmt.Println(0)
// 		return
// 	}

// 	oper := arg[1]
// 	var res int64

// 	switch oper {
// 	case "+":
// 		res = x + y
// 		if x < 0 && y < 0 && res > 0 {
// 			fmt.Println(0)
// 			return
// 		}
// 		if x > 0 && y > 0 && res < 0 {
// 			return
// 		}

// 	case "-":
// 		res = x - y
// 		if x < 0 && y > 0 && res > 0 {
// 			fmt.Println("0")
// 			return
// 		}
// 		if x > 0 && y < 0 && res < 0 {
// 			fmt.Println(0)
// 			return
// 		}
// 	case "*":
// 		res = x * y     // 15 = 5 * 3
// 		if res/x != y { // 15 / 5 = 3
// 			return
// 		}
// 	case "/":
// 		if arg[2] == "0" {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		res = x / y
// 		if x < 0 && y < 0 && res < 0 {
// 			fmt.Println(0)
// 			return
// 		}

// 	case "%":
// 		if arg[2] == "0" {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		res = x % y
// 	}
// 	fmt.Println(res)
// }

// // Atoi : conv from string to int with IsNumeric
// func Atoi(s string) (int64, bool) {
// 	var res int64 = 0
// 	neg := false
// 	if s[0] == '+' {
// 		s = s[1:]
// 	} else if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}
// 	if len(s) == 0 || len(s) > 19 {
// 		return 0, true
// 	}
// 	for _, el := range s {l
// 		}
// 	}
// 	if neg {
// 		res *= -1
// 	}
// 	if !neg && res < 0 {
// 		return 0, true
// 	}
// 	if neg && res > 0 {
// 		return 0, true
// 	}
// 	return res, false // - > если ошибки нет
// }

//*************************************************>>>ROT13

//***********************************************first var

//***************************************************lastrune
// import (
// 	"github.com/01-edu/z01"
// )

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }
// func LastRune(s string) rune {
// 	res := []rune(s) // first variant
// 	return res[len(s)-1]
// 	// return rune(s[len(s)-1]) // second variant
// }

//*******************************************************>>>FirstRune
// import (
// 	"github.com/01-edu/z01"

// 	"piscine"
// )

// func main() {
// 	z01.PrintRune(piscine.FirstRune("Hello!"))
// 	z01.PrintRune(piscine.FirstRune("Salut!"))
// 	z01.PrintRune(piscine.FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// func FirstRune(s string) rune {
// 	return rune(s[0])
// }

//*******************************************************second variant of firstrune
// func FirstRune(s string) rune {
// 	res := []rune(s) // changing s to []rune
// 	return res[0]    // return first rune by index examole for "Hello!": []rune{"Hello!"} -> idx of array [0=H,1=e,2=l,3=l,4=0]
// }

// ****************************************************>>>StrLen
// import (
// 	"fmt"
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }

// func StrLen(s string) int {
// 	return len([]rune(s))
// }
//*****************************************************>>>CountDown
// func main() {
// 	for i := 9; i >= 0; i-- {
// 		z01.PrintRune(rune(i + 48))
// 	}
// 	z01.PrintRune(10)
// }

// func main() {
// 	for i := '9'; i >= '0'; i-- {
// 		z01.PrintRune(rune(i))
// 	}
// 	z01.PrintRune(10)
// }

//**************************************************>>>Compare
// func main() {
// 	fmt.Println(Compare("Hello!", "Hello!"))
// 	fmt.Println(Compare("Salut!", "lut!"))
// 	fmt.Println(Compare("Ola!", "Ol"))
// }

// func Compare(a, b string) int {
// 	// if a == b {
// 	// 	return 0
// 	// }
// 	if a < b {
// 		return -1
// 	}
// 	if a > b {
// 		return 1
// 	}
// 	return 0
// }
