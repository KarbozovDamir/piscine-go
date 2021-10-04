package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) == 2 {
		word := ""
		lastWord := ""
		for _, el := range os.Args[1] + " " {
			if el == ' ' {
				if len(word) > 0 {
					lastWord = word
				}
				word = ""
			} else {
				word += string(el)
			}
		}
		if lastWord == "" {
			return
		}
		for _, j := range lastWord {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}

// func main() {
// 	if len(os.Args) == 2 {
// 		lastWord := ""
// 		for i := len(os.Args[1]) - 1; i >= 0; i-- {
// 			if os.Args[1][i] == ' ' {
// 				if lastWord != "" {
// 					break
// 				}
// 			} else {
// 				lastWord += string(os.Args[1][i])
// 			}
// 		}
// 		if lastWord == "" {
// 			return
// 		}
// 		for i := len(lastWord) - 1; i >= 0; i-- {
// 			z01.PrintRune(rune(lastWord[i]))
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune(10)
// }

// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 1 {
// 		return
// 	}
// 	a := -1
// 	for i := len(arg[0]) - 1; i > 0; i-- {
// 		if a != -1 && arg[0][i] == ' ' {
// 			fmt.Println(arg[0][i+1 : a+1])
// 			return
// 		}
// 		if arg[0][i] == ' ' {
// 			continue
// 		}
// 		if arg[0][i] != ' ' && a == -1 {
// 			fmt.Println(i)
// 			a = i
// 		}

// 	}
// }

// if len(os.Args) == 2 {
// 	isWord := false
// 	word := ""
// 	lastword := ""
// 	for _, el := range os.Args[1] {
// 		if el != ' ' {
// 			isWord = true
// 		} else {
// 			isWord = false
// 			word = ""
// 		}
// 		if isWord {
// 			word += string(el)
// 			lastword = word
// 		}
// 	}

// 	if word > "" {
// 		for _, el := range lastword {
// 			z01.PrintRune(el)
// 		}
// 		z01.PrintRune(10)
// 	} else if lastword > "" {
// 		for _, el := range lastword {
// 			z01.PrintRune(el)
// 		}
// 		z01.PrintRune(10)
// 	}
// }

//****************************************************lastwvar1
// func main() {
// 	if len(os.Args) == 2 {
// 		word := ""
// 		lastWord := ""
// 		for _, el := range os.Args[1] + " " {
// 			if el == ' ' {
// 				if len(word) > 0 {
// 					lastWord = word
// 				}
// 				word = ""
// 			} else {
// 				word += string(el)
// 			}
// 		}

// 		if lastWord == "" {
// 			return
// 		}
// 		for _, j := range word {
// 			z01.PrintRune(j)
// 		}
// 		z01.PrintRune(10)
// 	}
// }

//**********************************************************lastwvar2
// func main() {
// 	if len(os.Args) == 2 {
// 		arg := os.Args[1]
// 		if arg == "" {
// 			return
// 		}
// 		word := ""
// 		lastWord := ""
// 		for _, el := range arg + " " {
// 			if el == ' ' {
// 				if word != "" {
// 					lastWord = word
// 				}
// 				word = ""

// 			} else {
// 				word = string(el)
// 			}
// 		}
// 		// if word != "" {
// 		// 	lastWord = word
// 		// }
// 		for _, j := range lastWord {
// 			z01.PrintRune(j)
// 		}
// 		if lastWord != "" {
// 			z01.PrintRune(10)
// 		}
// 	}
// }
//*********************************************************lastwvar3
// if len(os.Args) == 2 {
// 	lastWord := ""
// 	for i := len(os.Args[1]) - 1; i >= 0; i-- {
// 		if os.Args[1][i] == ' ' {
// 			if lastWord != "" {
// 				break
// 			}
// 		} else {
// 			lastWord += string(os.Args[1][i])
// 		}
// 	}
// 	if lastWord == "" {
// 		return
// 	}
// 	for i := len(lastWord) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(lastWord[i]))
// 	}
// } else {
// 	return
// }
// z01.PrintRune(10)

//************************other var-s
// func main() {
// 	if len(os.Args) == 2 {
// 		isWord := false
// 		word := ""
// 		lastword := ""
// 		for _, el := range os.Args[1] {
// 			if el != ' ' {
// 				isWord = true
// 			} else {
// 				isWord = false
// 				word = ""
// 			}
// 			if isWord {
// 				word += string(el)
// 				lastword = word
// 			}
// 		}

// 		if word > "" {
// 			for _, el := range lastword {
// 				z01.PrintRune(el)
// 			}
// 			z01.PrintRune(10)
// 		} else if lastword > "" {
// 			for _, el := range lastword {
// 				z01.PrintRune(el)
// 			}
// 			z01.PrintRune(10)
// 		}
// 	}
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		word := ""
// 		/* lastWord := ""*/
// 		arr := []string{}
// 		for _, el := range os.Args[1] + " " {
// 			if el == ' ' {
// 				if len(word) > 0 {
// 					/*// lastWord = word*/
// 					arr = append(arr, word)
// 				}
// 				word = ""
// 			} else {
// 				word += string(el)
// 			}
// 		}
// 		/*// if lastWord == "" {
// 		// 	return
// 		// }
// 		// for _, el := range lastWord {
// 		// 	z01.PrintRune(el)
// 		// }*/
// 		if len(arr) < 2 {
// 			fmt.Println("pred no")
// 			return
// 		}
// 		for _, el := range arr[len(arr)-2] {
// 			z01.PrintRune(el)
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune(10)
// }
