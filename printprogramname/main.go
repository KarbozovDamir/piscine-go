package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := os.Args[2]
	for _, el := range result {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}
