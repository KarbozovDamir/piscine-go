package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune(10)
		return
	}
	if len(os.Args) == 3 {
		join := os.Args[1] + os.Args[2]
		mp := map[rune]int{}
		for _, el := range join {
			if mp[el] < 1 {
				mp[el]++
				z01.PrintRune(el)
			}
		}
		z01.PrintRune(10)
	}
}
