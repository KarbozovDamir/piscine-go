package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {

// 	res := ""
// 	if len(os.Args) != 2 {
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

func main() {
	if len(os.Args) != 2 {
		return
	}

	for _, el := range os.Args[1] {
		if el >= 'a' && el <= 'm' || el >= 'A' && el <= 'M' {
			z01.PrintRune(el + 13)
		} else if el > 'm' && el <= 'z' || el > 'M' && el <= 'Z' {
			z01.PrintRune(el - 13)
		} else {
			z01.PrintRune(el)
		}
	}
	z01.PrintRune(10)
}
