package main

import (
	"fmt"
	"os"
)

//*********************************Union
func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Print("\n")
		return
	}
	if len(os.Args[1:]) == 1 {
		fmt.Println("\n")
		return
	}
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		s1 += s2
		res := ""
		delete := make(map[rune]bool)
		for _, ch := range s1 {
			if delete[ch] == false {
				res += string(ch)
				delete[ch] = true
			}
		}
		fmt.Println(res)
	}
}

//***********************************************
// func Atoi(s string) (n int) {
// 	var res int = 0
// 	neg := false
// 	if s[0] == '+' {
// 		s = s[1:]
// 	}
// 	if s[0] == '-' {
// 		s = s[1:]
// 	}
// 	for _, el := range s {
// 		if el >= '0' && el <= '9' {
// 			res = res*10 + int(el-48)

// 		} else {
// 			return 0
// 		}
// 	}
// 	if neg {
// 		res *= -1
// 	}
// 	if !neg && res < 0 {
// 		return 0
// 	}
// 	if neg && res > 0 {
// 		return 0
// 	}
// 	return res
// }
//**************************************balancedstring
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) == 2 {
// 		C := 0 // it is var is not rune
// 		D := 0 // still var is not rune
// 		counter := 0
// 		for _, el := range os.Args[1] {
// 			if el == 'C' {
// 				C++
// 			} else {
// 				D++
// 			}
// 			if C == D {
// 				counter++
// 			}
// 		}
// 		fmt.Println(counter)
// 	}
// }

//************************************************TabMult
// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) == 2 {
// 		var res string
// 		n, _ := strconv.Atoi(os.Args[1])
// 		for i := 1; i <= 9; i++ {
// 			res += strconv.Itoa(i) + " x " + strconv.Itoa(n) + " = " + strconv.Itoa(n*i) + "\n"
// 		}
// 		for _, j := range res {
// 			z01.PrintRune(j)
// 		}

// 	}
// }
// include this Itoa
// func Itoa(n int) (res string) {
// 	if n > 0 {
// 		res = string(n%10+48) + res

// 		n /= 10
// 	}
// 	return
// }

//***********************************************INTER
// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $

// func main() {
// 	if len(os.Args) == 3 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		res := ""
// 		for _, el1 := range s1 {
// 			for _, el2 := range s2 {
// 				if el1 == el2 {
// 					isUniq := true
// 					for _, el3 := range res {
// 						if el3 == el1 {
// 							isUniq = false
// 						}
// 					}
// 					if isUniq {
// 						res += string(el1)
// 					}
// 				}
// 			}
// 		}
// 		os.Stdout.WriteString(res + "\n")
// 	}
// }

//**************************************************TwoSum
// func main() {
// 	case1 := []int{1, 2, 3, 4}
// 	out := TwoSum(case1, 5)
// 	fmt.Println(out)
// }

// func TwoSum(nums []int, target int) []int {
// 	for i := 0; i <= len(nums); i++ {
// 		for j := i + 1; j <= len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

/*
   rostring
   Instructions

   Write a program that takes a string and displays this string after rotating it one word to the left.

   Thus, the first word becomes the last, and others stay in the same order.

   A word is a sequence of alphanumerical characters.

   Words will be separated by only one space in the output.

   If the number of arguments is different from 1, the program displays a newline.
   Usagego run . "FOR PONY" | cat -e
*/
//******************************************************NAuuo
// func main() {
// 	fmt.Println(Nauuo(50, 43, 20))
// 	fmt.Println(Nauuo(13, 13, 0))
// 	fmt.Println(Nauuo(10, 9, 0))
// 	fmt.Println(Nauuo(5, 9, 2))
// }
// func Nauuo(plus, minus, rand int) string {
// 	if plus > minus+rand {
// 		return "+"
// 	} else if minus > plus+rand {
// 		return "-"
// 	} else if plus == minus && rand == 0 {
// 		return "0"
// 	} else {
// 		return "?"
// 	}
// }

//***********************************************2 variant
// func Nauuo(plus, minus, rand int) string {
// 	res := ""
// 	diferent := plus - minus
// 	if diferent < 0 {
// 		diferent = minus - plus
// 	}
// 	if diferent == 0 && rand == 0 {
// 		return "0"
// 	}
// 	if diferent <= rand {
// 		return "?"
// 	} else if plus == minus {
// 		return "0"
// 	} else if plus > minus {
// 		return "+"
// 	} else if plus < minus {
// 		return "-"
// 	}
// 	return res
// }

//*****************************ispowerof2
// func main() {
// 	if len(os.Args[1:]) != 1 {
// 		return
// 	}
// 	x, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		return
// 	}

// 	for x > 1 {
// 		if x%2 == 1 {
// 			fmt.Println("false")
// 			return
// 		}
// 		x /= 2
// 	}
// 	fmt.Println("true")
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

//*******************************************************doop
// import (
// 	"fmt"
// 	"os"
// )

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
// 	for _, el := range s {
// 		if el >= '0' && el <= '9' {
// 			res = res*10 + int64(el-48)
// 		} else {
// 			return 0, true
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

// Atoi : convert from string to int
// func atoi(s string) int {
// 	res := 0
// 	neg := false
// 	if s[0] == '+' {
// 		s = s[1:]
// 	}
// 	if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}
// 	for _, el := range s {
// 		if el >= '0' && el <= '9' {
// 			res = res*10 + int(el-48)
// 		} else {
// 			return 0
// 		}
// 	}
// 	if neg {
// 		res *= -1
// 	}
// 	if !neg && res < 0 {
// 		return 0
// 	}
// 	if neg && res > 0 {
// 		return 0
// 	}
// 	return res
// }

// func Checker(b string) bool {
// 	checker := true
// 	if b != "/" && b != "%" && b != "*" && b != "+" && b != "-" {
// 		checker = false
// 	}
// 	return checker
// }

//*********************************************Alphamirror
// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )
//********************************first variant
// func main() {
// 	if len(os.Args) == 2 {
// 		for _, j := range os.Args[1] {
// 			if j >= 'A' && j <= 'Z' {
// 				z01.PrintRune('Z' - j + 'A')
// 			} else if j >= 'a' && j <= 'z' {
// 				z01.PrintRune('z' - j + 'a')
// 			} else {
// 				z01.PrintRune(j)
// 			}
// 		}
// 		z01.PrintRune(10)
// 	}
// }
//*******************************second variant
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	args := os.Args[1:]

// 	for _, j := range args[0] {
// 		if j >= 'A' && j <= 'Z' {
// 			z01.PrintRune('Z' - j + 'A')

// 		} else if j >= 'a' && j <= 'z' {
// 			z01.PrintRune('z' - j + 'a')
// 		} else {
// 			z01.PrintRune(j)
// 		}
// 	}
// 	z01.PrintRune(10)
// }
//**********************************************************3 variant
// func main() {
// 	if len(os.Args) == 2 {
// 		res := ""
// 		for _, el := range os.Args[1] {
// 			if el >= 'a' && el <= 'z' {
// 				res += string('z' - el + 'a')
// 			} else if el >= 'A' && el <= 'Z' {
// 				res += string('Z' - el + 'A')
// 			} else {
// 				res += string(el)
// 			}
// 		}
// 		for _, el := range res {
// 			z01.PrintRune(el)
// 		}
// 	}
// 	z01.PrintRune(10)
// }

//***************************************************************lastword
// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )
// func main() {
// 	if len(os.Args) == 2 {
// 		isWord := false
// 		word := ""
// 		lastword := ""
// 		for _, el := range os.Args[1] {
// 			if el != ' ' {
// 				isWord = true
// 			} else {
// 				isWord = false
// 				word = ""
// 			}
// 			if isWord {
// 				word += string(el)
// 				lastword = word
// 			}
// 		}

// 		if word > "" {
// 			for _, el := range lastword {
// 				z01.PrintRune(el)
// 			}
// 			z01.PrintRune(10)
// 		} else if lastword > "" {
// 			for _, el := range lastword {
// 				z01.PrintRune(el)
// 			}
// 			z01.PrintRune(10)
// 		}
// 	}
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		word := ""
// 		/* lastWord := ""*/
// 		arr := []string{}
// 		for _, el := range os.Args[1] + " " {
// 			if el == ' ' {
// 				if len(word) > 0 {
// 					/*// lastWord = word*/
// 					arr = append(arr, word)
// 				}
// 				word = ""
// 			} else {
// 				word += string(el)
// 			}
// 		}
// 		/*// if lastWord == "" {
// 		// 	return
// 		// }
// 		// for _, el := range lastWord {
// 		// 	z01.PrintRune(el)
// 		// }*/
// 		if len(arr) < 2 {
// 			fmt.Println("pred no")
// 			return
// 		}
// 		for _, el := range arr[len(arr)-2] {
// 			z01.PrintRune(el)
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune(10)
// }

//*************************************************>>>ROT13
// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )
//***********************************************first var

// if len(os.Args) == 2 {
// 	res := ""
// 	for _, el := range os.Args[1] {
// 		if el >= 'A' && el < 'M' || el >= 'a' && el < 'm' {
// 			res += string(el + 13)
// 		} else if el > 'M' && el <= 'Z' || el > 'm' && el <= 'z' {
// 			res += string(el - 13)
// 		} else {
// 			res += string(el)
// 		}
// 	}
// 	for _, j := range res {

// 		z01.PrintRune(j)
// 	}
// }
// z01.PrintRune(10)
// }
//*************************************************************
// func main() {
// 	args := os.Args[1:]
// 	res := ""
// 	if len(args) != 1 {
// 		return
// 	}
// 	for _, el := range os.Args[1] {

// 		if el >= 'a' && el <= 'm' || el >= 'A' && el <= 'M' {
// 			res += string(el + 13)

// 		} else if el > 'm' && el <= 'z' || el > 'M' && el <= 'Z' {
// 			res += string(el - 13)
// 		} else {
// 			res += string(el)
// 		}
// 	}
// 	for _, el := range res {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

//*************************1 var rot13

// func main() {
// 	if len(os.Args) == 2 {
// 		for _, el := range os.Args[1] {
// 			if el >= 'a' && el < 'm' {
// 				z01.PrintRune(el + 13)
// 			} else if el > 'm' && el <= 'z' {
// 				z01.PrintRune(el - 13)
// 			} else if el >= 'A' && el < 'M' {
// 				z01.PrintRune(el + 13)
// 			} else if el > 'M' && el <= 'Z' {
// 				z01.PrintRune(el - 13)
// 			} else {
// 				z01.PrintRune(el)
// 			}
// 		}go run . "abc"
// }
//*****************************************Tabmult
// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func Itoa(number int) (result string) {
// 	for number > 0 {
// 		result = string(number%10+48) + result
// 		number /= 10
// 	}
// 	return
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		var result string
// 		number, _ := strconv.Atoi(os.Args[1])
// 		for i := 1; i <= 9; i++ {
// 			result += Itoa(i) + " x " + Itoa(number) + " = " + Itoa(number*i) + "\n"
// 		}
// 		for _, r := range result {
// 			z01.PrintRune(r)
// 		}
// 	}
// }

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
