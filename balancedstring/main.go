package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		C := 0 // it is var is not rune
		D := 0 // still var is not rune
		counter := 0
		for _, el := range os.Args[1] {
			if el == 'C' {
				C++
			} else {
				D++
			}
			if C == D {
				counter++
			}
		}
		fmt.Println(counter)
	}
}

// func main() {
// 	args := os.Args[1:]

// 	countC := 0
// 	countD := 0
// 	counter := 0
// 	for _, el := range args[0] {
// 		if el == 'C' {
// 			countC++
// 		} else if el == 'D' {
// 			countD++
// 		}
// 		if countC == countD {
// 			counter++
// 		}
// 	}
// 	n := itoa(counter)
// 	for i := len(n) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(n[i]))
// 	}
// 	z01.PrintRune('\n')
// }

// func itoa(i int) string {
// 	str := ""
// 	for i > 0 {
// 		str += string(i%10 + 48)
// 		i /= 10
// 	}
// 	return str
// }
