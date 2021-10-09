package main

import (
	"github.com/01-edu/z01"
)

func main() {
	s := "Hello World!"
	for _, el := range s {
		z01.PrintRune(el)
	}
	z01.PrintRune(10)
}
