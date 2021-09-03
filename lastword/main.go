package main

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
