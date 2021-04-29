package main

import "github.com/01-edu/z01"

func main() {
	var letter string = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i <= 25; i++ {
		z01.PrintRune(rune(letter[i]))
	}
	z01.PrintRune('\n')
}
