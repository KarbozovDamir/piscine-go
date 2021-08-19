package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func Itoa(number int) (result string) {
	for number > 0 {
		result += string(number%10 + 48)
		number /= 10
	}
	return
}

func main() {
	if len(os.Args) == 2 {
		var result string
		number, _ := strconv.Atoi(os.Args[1])
		for i := 1; i <= 9; i++ {
			result += Itoa(i) + " x " + Itoa(number) + " = " + Itoa(number*i) + "\n"
		}
		for _, r := range result {
			z01.PrintRune(r)
		}
	}
}

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
//*************************************************>>>ROT13

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
// 			// fmt.Println("")
// 		}
// 	}
// 	fmt.Println(res)
// }
//***************************************************************

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	}
// 	s1 := args[0]
// 	s2 := args[1]
// 	res := ""
// 	a := true
// 	for _, el := range s1 {
// 		for _, j := range s2 {
// 			if el == j {
// 				a = false
// 				break
// 			}
// 		}
// 		for _, k := range res {
// 			if k == el {
// 				a = true
// 			}
// 		}
// 		if !a {
// 			res += string(el)
// 		}
// 		a = true
// 	}
// 	fmt.Println(res)
// }
