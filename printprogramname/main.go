package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, letter := range os.Args[0] {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
