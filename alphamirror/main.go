package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, j := range os.Args[1] {
			if j >= 'A' && j <= 'Z' {
				z01.PrintRune('Z' - j + 'A')
			} else if j >= 'a' && j <= 'z' {
				z01.PrintRune('z' - j + 'a')
			} else {
				z01.PrintRune(j)
			}
		}
		z01.PrintRune(10)
	}
}
