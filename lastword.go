package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	args := os.Args[1:]
// 	// s := []rune(args[0])
// 	s := args[0]

// 	fmt.Println(s)
// }

func main() {
	if len(os.Args) == 2 {
		isWord := false
		word := ""
		wordSlice := []string{}
		lastword := ""
		for _, el := range os.Args[1] {
			if el != ' ' {
				isWord = true
			} else {
				isWord = false
				if word != "" {
					wordSlice = append(wordSlice, word)
					word = ""
				}
			}
			if isWord {
				word += string(el)
			}
		}
		if word != "" {
			for _, el := range word {
				z01.PrintRune(el)
			}
		} else {
			for _, el := range wordSlice[len(wordSlice)-1] {
				z01.PrintRune(el)
			}
		}
		// if word > "" {
		// 	for _, el := range word {
		// 		z01.PrintRune(el)
		// 	}
		// 	z01.PrintRune('\n')
		// } else if lastword > "" {
		// 	for _, el := range lastword {
		// 		z01.PrintRune(el)
		// 	}
		// 	z01.PrintRune('\n')
		// }
	}
}
