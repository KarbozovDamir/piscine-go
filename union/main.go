package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) != 2 {
		fmt.Println(0)
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	s1 += s2
	res := ""
	delete := make(map[rune]bool)
	for _, ch := range s1 {
		if delete[ch] == false {
			res += string(ch)
			delete[ch] = true
		}
	}
	fmt.Println(res)
}

// func main() {
// 	if len(os.Args) == 3 {
// 		str1 := os.Args[1]
// 		str2 := os.Args[2]
// 		result := ""
// 		for _, a := range str1 {
// 			isUnique := true
// 			for _, b := range result {
// 				if a == b {
// 					isUnique = false
// 				}
// 			}
// 			if isUnique {
// 				result += string(a)
// 			}
// 		}
// 		for _, a := range str2 {
// 			isUnique := true
// 			for _, b := range result {
// 				if a == b {
// 					isUnique = false
// 				}
// 			}
// 			if isUnique {
// 				result += string(a)
// 			}
// 		}
// 		os.Stdout.WriteString(result + "\n")
// 	} else {
// 		os.Stdout.WriteString("\n")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	if len(os.Args) == 3 {
// 		var res string
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]

// 		for _, v := range s1 {
// 			if !strings.ContainsRune(res, v) {
// 				res += string(v)
// 			}
// 		}
// 		for _, v := range s2 {
// 			if !strings.ContainsRune(res, v) {
// 				res += string(v)
// 			}
// 		}
// 		fmt.Print(res)
// 	}
// 	fmt.Println()
// }

// func Union(args []string) {
// 	args = args[1:]
// 	if len(args) != 2 {
// 		return
// 	}
// 	uniques := getUniques(args)
// 	printUniques(uniques, args)
// }

// //
// func getUniques(strs []string) map[rune]bool {
// 	uniques := make(map[rune]bool)
// 	for _, word := range strs {
// 		for _, run := range word {
// 			uniques[run] = true
// 		}
// 	}
// 	return uniques
// }

// func printUniques(uniques map[rune]bool, strs []string) {
// 	for _, word := range strs {
// 		for _, run := range word {
// 			if uniques[run] {
// 				fmt.Printf("%v", string(run))
// 				uniques[run] = false
// 			}
// 		}
// 	}
// 	fmt.Println()
// }
// $ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
// zpadintoqefwjy$
// $ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
// df6vewg4thras$
// $ go run . rien "cette phrase ne cache rien" | cat -e
// rienct phas$
// $ go run . | cat -e
// $
// $ go run . rien | cat -e
// $
// $
