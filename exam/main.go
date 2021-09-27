package main

// *********************************************Split
// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Printf("%#v\n", Split(s, "HA"))
// }

// func Split(s, sep string) []string {
// 	prev := 0 // H_..o
// 	ans := make([]string, 0)
// 	for i := 0; i < len(s)-len(sep); i++ {
// 		if s[i:i+len(sep)] == sep {
// 			ans = append(ans, s[prev:i])
// 			i += len(sep) // for skip HA
// 			prev = i      // FOR START AFTER SKIP HA
// 		}
// 	}
// 	if prev != len(s)-1 {
// 		ans = append(ans, s[prev:])
// 	}
// 	return ans
// }

//*************************************chess
// func main() {
// 	if len(os.Args[1:]) != 2 {
// 		return
// 	}
// func Split(s, sep string) []string {

// }
// 	for i := 1; i <= n1; i++ {
// 		for j := 1; j <= n2; j++ {
// 			if i%2 == 1 && j%2 == 1 {
// 				z01.PrintRune('#')
// 			} else if i%2 == 1 && j%2 == 0 {
// 				z01.PrintRune(' ')
// 			} else if i%2 == 0 && j%2 == 1 {
// 				z01.PrintRune(' ')
// 			} else {
// 				z01.PrintRune('#')
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

//****************************************hiddenp
// counter := 0
// 	if len(os.Args) != 3 {
// 		returngithub.com/01-edu/go-tests/lib/challenge":= range os.Args[2] {
// 		if len(os.Args[1]) == counter {
// 			z01.PrintRune('1')
// 			z01.PrintRune(10)
// 			return
// 		}
// 		if v == rune(os.Args[1][counter]) {
// 			counter++
// 		}
// 	}
// 	z01.PrintRune('0')
// 	z01.PrintRune(10)
//*************************************hex
// func main() {
// 	if len(os.Args) == 2 {
// 		base := "0123456789abcdef"
// 		res := ""
// 		num, err := strconv.Atoi(os.Args[1])
// 		if err != nil {
// 			os.Exit(0)
// 		}
// 		for num != 0 {
// 			res = string(base[num%16]) + res
// 			num /= 16
// 		}
// 		if res == "" {
// 			os.Stdout.WriteString("ERROR")
// 		}
// 		os.Stdout.WriteString(res + "\n")
// 	}
// }

//***********************inter duplicate
/*func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		arr := []rune{}
		for _, x := range args[0] {
			for y := 0; y < len(args[1]); y++ {
				if byte(x) == args[1][y] && isIn(x, arr) {
					arr = append(arr, x)
				}
			}
		}
		for _, x := range arr {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
	// return
}

func isIn(char rune, arr []rune) bool {
	for _, x := range arr {
		if x == char {
			return false
		}
	}
	return true
}
*/
//**********************************repeatAlpha
// func main() {

// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	args := os.Args[1]
// 	if args == "" {
// 		z01.PrintRune(10)
// 		return
// 	}
// 	for _, el := range args {
// 		if el >= 'A' && el <= 'Z' {
// 			for i := 0; i <= int(el)-'A'; i++ {
// 				z01.PrintRune(el)
// 			}
// 		}
// 		if el >= 'a' && el <= 'z' {
// 			for i := 0; i <= int(el)-'a'; i++ {
// 				z01.PrintRune(el)
// 	if len(os.Args) == 3 {
// 		mp := map[rune]int{}
// 		for _, el := range os.Args[1] + os.Args[2] {
// 			if mp[el] < 1 {
// 				mp[el]++
// 				z01.PrintRune(el)
// 			}
// 		}
// 	}
// 	z01.PrintRune(10)
// }

//***********************************************************************printchessboard
// func main() {
// 	if len(os.Args) != 3 {
// 		error()
// 		return
// 	}
// 	x, err := strconv.Atoi(os.Args[1])
// 	y, err2 := strconv.Atoi(os.Args[2])

// 	if err != nil || err2 != nil {
// 		error()
// 		return
// 	}
// 	if x == 0 || y == 0 {
// 		error()
// 		returnfor _, el := range flag {
// 	z01.PrintRune(el)
// }
// z01.PrintRune(10)
// return
// 		for j := 1; j <= x; j++ {
// 			if (i+j)%2 == 0 {
// 				z01.PrintRune('#')
// 			} else {
// 				z01.PrintRune(' ')
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func error() {
// 	for _, el := range "Error" {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

//*******************************************************hiddenp

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return
// 	}
// 	count := 0
// 	i := 0
// 	for _, el := range arg[1] {
// 		if rune(arg[0][i]) == el {
// 			i++
// 			if i == len(arg[0]) {
// 				i = 0
// 				count++
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }

///****************************************isPowerOf2
// func main() {
// 	if len(os.Args) == 2 {
// 		n, _ := strconv.Atoi(os.Args[1])
// 		flag := "true"
// 		for n > 1 {
// 			if n%2 != 0 {
// 				flag = "false"
// 				n /= 2
// 			}
// for _, el := range flag {
// 	z01.PrintRune(el)
// }
// z01.PrintRune(10)
// return
// 		}
// 	}
// }

//*******************************************************hiddenp
// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return
// 	}
// 	count := 0
// 	i := 0
// 	for _, el := range arg[1] {
// 		if rune(arg[0][i]) == el {
// 			i++
// 			if i == len(arg[0]) {
// 				i = 0
// 				count++
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }
//*******************************************************hiddenp
// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return
// 	}
// 	count := 0
// 	i := 0
// 	for _, el := range arg[1] {
// 		if rune(arg[0][i]) == el {
// 			i++
// 			if i == len(arg[0]) {
// 				i = 0
// 				count++
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }

//*********************************************UNION Alibek
// 	if len(os.Args[1:]) == 0 {
// 		fmt.Print("\n")
// 		return
// 	}
// 	if len(os.Args[1:]) == 1 {
// 		fmt.Println()
// 		return
// 	}
// 	if len(os.Args) == 3 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		s1 += s2
// 		res := ""
// 		delete := make(map[rune]bool)
// 		for _, ch := range s1 {
// 			if delete[ch] == false {
// 				res += string(ch)
// 				delete[ch] = true
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

//****************************************************lastwvar1
// func main() {
// 	if len(os.Args) == 2 {
// 		word := ""
// 		lastWord := ""
// 		for _, el := range os.Args[1] + " " {
// 			if el == ' ' {
// 				if len(word) > 0 {
// 					lastWord = word
// 				}
// 				word = ""
// 			} else {
// 				word += string(el)
// 			}
// 		}

// 		if lastWord == "" {
// 			return
// 		}
// 		for _, j := range word {
// 			z01.PrintRune(j)
// 		}
// 		z01.PrintRune(10)
// 	}
// }

//**********************************************************lastwvar2
// func main() {
// 	if len(os.Args) == 2 {
// 		arg := os.Args[1]
// 		if arg == "" {
// 			return
// 		}
// 		word := ""
// 		lastWord := ""
// 		for _, el := range arg + " " {
// 			if el == ' ' {
// 				if word != "" {
// 					lastWord = word
// 				}
// 				word = ""

// 			} else {
// 				word = string(el)
// 			}
// 		}
// 		// if word != "" {
// 		// 	lastWord = word
// 		// }
// 		for _, j := range lastWord {
// 			z01.PrintRune(j)
// 		}
// 		if lastWord != "" {
// 			z01.PrintRune(10)
// 		}
// 	}
// }
//*********************************************************lastwvar3
// if len(os.Args) == 2 {
// 	lastWord := ""
// 	for i := len(os.Args[1]) - 1; i >= 0; i-- {
// 		if os.Args[1][i] == ' ' {
// 			if lastWord != "" {
// 				break
// 			}
// 		} else {
// 			lastWord += string(os.Args[1][i])
// 		}
// 	}
// 	if lastWord == "" {
// 		return
// 	}
// 	for i := len(lastWord) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(lastWord[i]))
// 	}
// } else {
// 	return
// }
// z01.PrintRune(10)

//**************************************doop
// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	x, err1 := Atoi(os.Args[1])
// 	oper := os.Args[2]
// 	y, err2 := Atoi(os.Args[3])
// 	res := 0
// 	if err1 == false || err2 == false {
// 		return
// 	}
// 	switch oper {
// 	case "+":
// 		res = x + y
// 		if x > 0 && res < 0 {
// 			return
// 		}
// 	case "-":
// 		res = x - y
// 		if x < 0 && res > 0 {
// 			return
// 		}
// 	case "*":
// 		res = x * y
// 		if res/x != y {
// 			return
// 		}
// 	case "/":
// 		if y == 0 {
// 			fmt.Println("No div by 0")
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

// // Atoi for convert from string to int with numeric
// func Atoi(s string) (int, bool) {
// 	var res int
// 	neg := false
// 	if len(s) == 0 {
// 		return 0, false
// 	} else if s[0] == '-' {
// 		s = s[1:]
// 		neg = true
// 	}

// 	for _, el := range s {
// 		if el < '0' || el > '9' {

// 			return 0, false
// 		}
// 		res = res*10 + (int(el) - 48)
// 	}
// 	if neg {
// 		res *= -1
// 	}
// 	return res, true
// }

//***************************************************GCD
// func main() {
// 		num2, _ := strconv.Atoi(os.Args[2])
// 		for i := num1; i > 0; i-- {
// 			if num1%i == 0 && num2%i == 0 {
// 				res = i
// 				break
// 			}
// 		}// func main() {
// 	if len(os.Args) == 2 {
// 		str := os.Args[1]
// 		for _, r := range str {
// 			if r >= 'A' && r <= 'Z' {
// 				for i := 0; i <= int(r)-'A'; i++ {
// 					z01.PrintRune(r)
// 				}
// 			}
// 			if r >= 'a' && r <= 'z' {
// 				for i := 0; i <= int(r)-'a'; i++ {
// 					z01.PrintRune(r)
// 				}
// 			} else {
// 				z01.PrintRune(r)
// 			}
// 		}
// 	}
// }
// 		return
// 	}
// }

//*****************************************lCM
// func Lcm(first, second int) int {
// 	res := 0
// 	for i := second; i <= first*second; i++ {
// 		if i%first == 0 && i%second == 0 {
// 			res = i
// 			break
// 		}
// 	}
// 	return res
// }

//************************************************************Repeatalpha
// func main() {
// 	if len(os.Args) == 2 {
// 		str := os.Args[1]
// 		for _, r := range str {
// 			if r >= 'A' && r <= 'Z' {
// 				for i := 0; i <= int(r)-'A'; i++ {
// 					z01.PrintRune(r)
// 				}
// 			}
// 			if r >= 'a' && r <= 'z' {
// 				for i := 0; i <= int(r)-'a'; i++ {
// 					z01.PrintRune(r)
// 				}
// 			} else {
// 				z01.PrintRune(r)
// 			}
// 		}
// 	}
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
//

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
/*
$ go run .
$ go run . 1 + 1 | cat -e
2$
$ go run . hello + 1
$ go run . 1 p 1
$ go run . 1 / 0 | cat -e
No division by 0$
$ go run . 1 % 0 | cat -e
No modulo by 0$
$ go run . 9223372036854775807 + 1
$ go run . -9223372036854775809 - 3
$ go run . 9223372036854775807 "*" 3
$ go run . 1 "*" 1
1
$ go run . 1 "*" -1
-1
$
*/

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
