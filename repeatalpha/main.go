package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Firstrune : fr
// func FirstRune(s string) rune {
// 	res := []rune(s)
// 	return res[0]
// }
// func FirstRune(s string) rune {
// 	return rune(s[0])
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}

func main() {
	if len(os.Args) != 2 {
		return
	}
	if len(os.Args) == 2 {
		for _, el := range os.Args[1] {

			if el >= 'A' && el <= 'Z' {
				for i := 'A'; i <= el; i++ {
					z01.PrintRune(el)
				}
			} else if el >= 'a' && el <= 'z' {
				for i := 'a'; i <= el; i++ {
					z01.PrintRune(el)
				}
			} else {
				z01.PrintRune(el)
			}
		}
		z01.PrintRune(10)
	}
}

//**************************
// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	for _, el := range os.Args[1] {
// 		if el >= 'a' && el <= 'z' {
// 			for i := 'a'; i <= el; i++ {
// 				z01.PrintRune(el)
// 			}
// 		} else if el >= 'A' && el <= 'Z' {
// 			for i := 'A'; i <= el; i++ {
// 				z01.PrintRune(el)
// 			}
// 		} else {
// 			z01.PrintRune(el)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	for _, x := range os.Args[1] {
// 		if x >= 65 && x <= 90 { // 65-A  90-Z
// 			for i := 1; i <= int(x-64); i++ {
// 				fmt.Print(string(x))
// 			}
// 		} else if x >= 97 && x <= 122 { // 97 -a 122 - z
// 			for i := 1; i <= int(x-96); i++ {
// 				fmt.Print(string(x))
// 			}
// 		} else {
// 			fmt.Print(string(x))
// 		}
// 	}
// 	fmt.Println()
// }

//*****************************
// func main() {
// 	if len(os.Args) == 2 {
// 		answer := ""
// 		for _, value := range os.Args[1] {
// 			switch {
// 			case value >= 'A' && value <= 'Z':
// 				dif := int(value + 1 - 'A')
// 				for i := 0; i < dif; i++ {
// 					answer += string(value)
// 				}
// 			case value >= 'a' && value <= 'z':
// 				dif := int(value + 1 - 'a')
// 				for i := 0; i < dif; i++ {
// 					answer += string(value)
// 				}
// 			default:
// 				answer += string(value)
// 			}
// 		}
// 		for _, runa := range answer {
// 			z01.PrintRune(runa)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

//*******************************************
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	for _, x := range args[0] {
// 		if x >= 'A' && x <= 'Z' {
// 			for i := 1; i <= int(x-64); i++ {
// 				z01.PrintRune(rune(x))
// 			}
// 		} else if x >= 97 && x <= 122 {
// 			for i := 1; i <= int(x-96); i++ {
// 				z01.PrintRune(rune(x))
// 			}
// 		} else {
// 			z01.PrintRune((x))
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
