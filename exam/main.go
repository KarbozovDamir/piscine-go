package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)


func main() {
if len(os.Args)
}

//hidden
// func main() {
// 	if len(os.Args) == 3 {
// 		arg1 := os.Args[1]
// 		arg2 := os.Args[2]
// 		i := 0
// 		if arg1 == arg2 {
// 			z01.PrintRune('1')
// 			z01.PrintRune(10)
// 			return
// 		}
// 		for j := range arg2 {
// 			if arg1[i] == arg2[j] {
// 				i++
// 				if i == len(arg1) {
// 					z01.PrintRune('1')
// 					z01.PrintRune(10)
// 					return
// 				}
// 			}
// 		}
// 		z01.PrintRune('0')
// 		z01.PrintRune(10)
// 	}
// }
//***********************************doop
// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	x, err1 := strconv.Atoi(os.Args[1])
// 	oper := os.Args[2]
// 	y, err2 := strconv.Atoi(os.Args[3])
// 	if err1 || err2 {
// 		return
// 	}
// 	res := 0
// 	switch oper {
// 	case "+":
// 		res = x + y
// 		if x > 0 && y > 0 && res < 0 {
// 			return
// 		}
// 		if x < 0 && y < 0 && res > 0 {
// 			return
// 		}
// 	case "-":
// 		res = x - y
// 		if x < y && res > 0 {
// 			return
// 		}
// 		if x > 0 && y < 0 && res < 0 {
// 			return
// 		}
// 	case "*":
// 		res = x * y
// 		if res/x != y {
// 			return
// 		}
// 	case "/":
// 		if y == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 	case "%":
// 		if y == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 	default:
// 		return
// 	}
// 	fmt.Println(res)

// }
// func Atoi(s string) (int, bool) {
// 	var res int
// 	neg := false
// 	if len(s) == 0 {
// 		return 0, false
// 	}
// 	if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}
// 	for _, el := range s {
// 		if el < '0' || el > '0' {
// 			return 0, true
// 		}
// 		if neg {
// 			res = res*10 - (int(el) - 48)
// 	}else {
// 		res = res*10 + (int(el) - 48)
// 		}

// 	if !neg && res < 0 || neg && res < 0 {
// 		return 0, true
// 	}
// 	return res, false
// }


// array with slize
// func main() {
// 	a := [3]int{1, 2, 3}

// 	// b := a
// 	s(a)
// 	// b[0] = 0
// 	fmt.Println(a) //1 2 3
// }

// func s(b [3]int) {
// 	b[0] = 5
// }

// 	a := [3]int{1, 2, 3}

// 	b := a

// 	b[0] = 0
// 	fmt.Println(a) //1 2 3
// }
*/

// func main() {
// 	fmt.Println(ispower(1))
// 	fmt.Println(ispower(2))
// 	fmt.Println(ispower(64))
// 	fmt.Println(ispower(255))
// }

// func ispower(num int) bool {
// 	if num == 1 {
// 		return true
// 	}
// 	if num == 0 {
// 		return false
// 	}
// 	if num%2 == 1 {
// 		return false
// 	}
// 	return ispower(num / 2)
// }

// DOOOOOOOOOOOOOOOOOOOOOOOOOOOOOP
// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}

// 	x, err1 := Atoi(os.Args[1])
// 	oper := os.Args[2]
// 	y, err2 := Atoi(os.Args[3])

// 	if !err1 || !err2 {
// 		return
// 	}
// 	res := 0
// 	switch oper {
// 	case "+":
// 		res = x + y
// 		if x > 0 && y > 0 && res < 0 { // 2 + 2 = -4 // res=10, n2=8, n1=?3
// 			return
// 		}

// 	case "-":
// 		res = x - y
// 		if x < y && res > 0 { // 8 - 1 = 7
// 			return
// 		}
// 	case "*":
// 		res = x * y

// 		if res/x != ystrconv.
// 	case "/":
// 		if y == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		res = x / y
// 	case "%":
// 		if y == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		res = x % y
// 	default:
// 		return
// 	}
// 	fmt.Println(res)
// }

// func Atoi(s string) (int, bool) {
// 	neg := false
// 	if len(s) == 0 {
// 		return 0, false
// 	}
// 	if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}

// 	var res int
// 	for _, el := range s {
// 		if el < '0' || el > '9' {
// 			return 0, false
// 		}
// 		res = res*10 + (int(el) - 48)
// 	}

// 	if neg {
// 		res *= -1
// 	}
// 	if !neg && res < 0 || neg && res > 0 {
// 		return 0, false
// 	}
// 	return res, true
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
