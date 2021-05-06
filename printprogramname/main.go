package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, el := range os.Args[0] {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
