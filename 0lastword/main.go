package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		isWord := false
		word := ""
		lastword := ""
		for _, el := range os.Args[1] {
			if el != ' ' {
				isWord = true
			} else {
				isWord = false
				word = ""
			}
			if isWord {
				word += string(el)
				lastword = word
			}
		}

		if word > "" {
			for _, el := range lastword {
				z01.PrintRune(el)
			}
			z01.PrintRune(10)
		} else if lastword > "" {
			for _, el := range lastword {
				z01.PrintRune(el)
			}
			z01.PrintRune(10)
		}
	}
}
