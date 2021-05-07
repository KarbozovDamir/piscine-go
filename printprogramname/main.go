package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	res := os.Args[0]
	for _, el := range res {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
