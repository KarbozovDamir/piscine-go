package main

import (
	"github.com/01-edu/z01"
)

func main() {

	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(-9223372036854775808)
	z01.PrintRune('\n')
}

// PrintNbr : nbr
func PrintNbr(n int) {
	num := []byte{}

	if n < 0 {
		z01.PrintRune('-')

	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 { // it is meaning until n changed to zero
		lastNum := n % 10
		if n < 0 {
			lastNum = -lastNum
		}
		num = append(num, byte(lastNum+48))
		n /= 10
	}
	for i := len(num) - 1; i >= 0; i-- {
		z01.PrintRune(rune(num[i]))

	}
}
