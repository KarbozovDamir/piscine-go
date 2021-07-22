package piscine

import (
	"github.com/01-edu/z01"
)

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

// func PrintNbr(n int) {
// 	if n < 0 {
// 		z01.PrintRune('-')
// 	}
// 	SetNbr(n)
// }

// func SetNbr(n int) {
// 	a := '0'
// 	if n == 0 {
// 		z01.PrintRune(a)
// 		return
// 	}
// 	for i := 1; i <= n%10; i++ {
// 		a++
// 	}
// 	for i := -1; i >= n%10; i-- {
// 		a++
// 	}
// 	if n/10 != 0 {
// 		SetNbr(n / 10)
// 	}
// 	z01.PrintRune(a)
// 	return
// }

// func print_next(n, dot int) {
// 	if n == 0 {
// 		return
// 	}
// 	digit := int((n % 10) * dot)
// 	_n := int(n / 10)
// 	print_next(_n, dot)
// 	z01.PrintRune(48 + rune(digit))
// }

// func PrintNbr(n int) {
// 	dot := 1
// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if n < 0 {
// 		dot *= -1
// 		z01.PrintRune('-')
// 	}
// 	print_next(n, dot)
// }
