package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}

func SetNbr(n int) {
	a := '0'
	if n == 0 {
		z01.PrintRune(a)
		return
	}
	for i := 1; i <= n%10; i++ {
		a++
	}
	for i := -1; i >= n%10; i-- {
		a++
	}
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	z01.PrintRune(a)
	return
}

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
