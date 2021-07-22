package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

// PrintNbr : nbr
func PrintNbr(n int) {

	num := []byte{}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n == 0 {
		z01.PrintRune('0')
	}

	for n != 0 {
		v := n % 10
		num = append(num, byte(v+48))
		n = n / 10
	}

	for i := len(num) - 1; i >= 0; i-- {
		z01.PrintRune(rune(num[i]))
	}
}
