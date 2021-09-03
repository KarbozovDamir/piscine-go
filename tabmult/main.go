package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	for i := 1; i <= 9; i++ {
		print(i, args)
		z01.PrintRune(10)
	}
}
func print(i int, args string) {
	z01.PrintRune(rune(i + '0'))
	z01.PrintRune(' ')

	z01.PrintRune('x')
	z01.PrintRune(' ')

	num, _ := strconv.Atoi(args)
	res := strconv.Itoa(i * num)

	for _, s := range strconv.Itoa(num) {
		z01.PrintRune(s) //вывод num, потому что z01 выводит только один символ
	}
	z01.PrintRune(' ')

	z01.PrintRune('=')
	z01.PrintRune(' ')

	for _, s := range res {
		z01.PrintRune(s)
	}
}
