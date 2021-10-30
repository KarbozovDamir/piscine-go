package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args[1:]) != 1 {
		fmt.Println()
		return
	}
	arg := os.Args[1]
	if arg == "" {
		fmt.Println()
		return
	}
	words := make([]string, 0)
	temp := ""

	for _, el := range arg{
		if el != ' ' {
			temp += string(el)
		}else {
			if temp != "" {
			words = append(words, temp)
			temp = ""
			}
		}
	
	}
	if temp != "" {
		words = append(words, temp)
		temp = ""
	} 
	words = append(words[1:], words[0])
	// fmt.Println(words)
	for i, word := range words {
		temp+=word
		if i != len(words) - 1 {
			temp += " "
		}
	}
	fmt.Println(temp)
	//["ali", "alibek", "damir"]
}
/*
Write a program that takes a string and displays this string after rotating it one word to the left.

Thus, the first word becomes the last, and others stay in the same order.

A word is a sequence of alphanumerical characters.

Words will be separated by only one space in the output.

If the number of arguments is different from 1, the program displays a newline.
*/
// func main() {
// 	arg := os.Args
// 	if len(arg) == 2 {
// 		ans := ""
// 		ff := ""
// 		tmp := ""
// 		arg[0] += " "
// 		for _, c := range arg[0] {
// 			if c == ' ' && tmp != "" {
// 				if ff == "" {
// 					ff = tmp

// 					tmp = ""
// 					continue
// 				}
// 				if ans == "" {
// 					ans = tmp
// 				} else {
// 					ans = ans + string(" ") + tmp
// 				}
// 				tmp = ""
// 			} else if c != ' ' {
// 				tmp = tmp + string(c)
// 			}
// 		}
// 		if ans == "" {
// 			ans = ff
// 		} else {
// 			ans = ans + string(" ") + ff
// 		}
// 		for _, c := range ans {
// 			z01.PrintRune(c)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }


// func main() {
// 	defer z01.PrintRune('\n')
// 	arg := os.Args
// 	if len(arg) != 2 {
// 		return
// 	}
// 	printStr(rostring(spacesBegin(spacesEnd(arg[0]))))
// }

// func printStr(s string) {
// 	for _, v := range s {
// 		z01.PrintRune(v)
// 	}
// }

// func spacesBegin(s string) string {
// 	for i, _ := range s {
// 		if s[0] == ' ' {
// 			if s[i] == ' ' && s[i+1] != ' ' {
// 				return s[i+1:]
// 			}
// 		}
// 	}
// 	return s
// }

// func spacesEnd(s string) string {
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[len(s)-1] == ' ' {
// 			if s[i] == ' ' && s[i-1] != ' ' {
// 				return s[:i]
// 			}
// 		}
// 	}
// 	return s
// }

// func rostring(s string) string {
// 	for i, _ := range s {
// 		if s[i] == ' ' && s[i+1] != ' ' {
// 			return s[i+1:] + " " + s[:i]
// 		}
// 	}
// 	return s
// }

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
