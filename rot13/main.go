package main

import (
	"os"

	"github.com/01-edu/z01"
)

//Rot13 : rot13
func main() {
	ars := os.Args[1:]
	s := []rune(ars[0])
	res := ""
	for _, j := range s {
		if j >= 'a' && j <= 'z' {
			if j >= 'n' {
				res += string(j - 13)
			} else {
				res += string(j + 13)
			}
		} else if j >= 'A' && j <= 'Z' {
			if j >= 'N' {
				res += string(j - 13)
			} else {
				res += string(j + 13)
			}
		} else {
			res += string(j)
		}
	}
	for _, h := range res {
		z01.PrintRune(h)
	}
	z01.PrintRune('\n')
}

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
