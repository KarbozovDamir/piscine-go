package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0])
	for _, res := range os.Args[0] {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
