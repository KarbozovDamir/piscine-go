package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0]) 
	for _, el := range runes {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
