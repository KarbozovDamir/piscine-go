package main

// import "fmt"

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	piscine.PrintNbr(-123)
	piscine.PrintNbr(0)
	piscine.PrintNbr(123)
	z01.PrintRune('\n')
}

// func main() {
// 	piscine.PrintComb2()
// }

func PrintNbr(n int) {
	p := 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	r := []int{}
	if n < 0 {
		z01.PrintRune('-')
		p = -1
	}
	for n != 0 {
		r = append(r, p*(n%10)+48)
		n /= 10
	}
	for i := len(r) - 1; i >= 0; i-- {
		z01.PrintRune(rune(r[i]))
	}
}

// func main() {
// 	for i := 0; i < 5; i++ {
// 		for j := 5 - i; j > 0; j-- {
// 			fmt.Print("*")
// 		}

// // PrintNbr : nbr
// func PrintNbr(n int) {
// 	num := []byte{}

// 	if n < 0 {
// 		z01.PrintRune('-')

// 	}
// 	if n == 0 {
// 		z01.PrintRune('0')
// 	}
// 	for n != 0 { // it is meaning until n changed to zero
// 		lastNum := n % 10
// 		if n < 0 {
// 			lastNum = -lastNum
// 		}
// 		num = append(num, byte(lastNum+48))
// 		n /= 10
// 	}
// 	for i := len(num) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(num[i]))

// 	}
// }
