package main

import (
	"os"
)

func main() {

	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		res := ""
		for _, a := range str1 {
			for _, b := range str2 {
				if a == b {
					isUnique := true
					for _, c := range res {
						if a == c {
							isUnique = false
						}
					}
					if isUnique {
						res += string(a)
					}
				}
			}
		}
		os.Stdout.WriteString(res + "\n")
	}
}

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $
//********************************Inter
// func main() {
// 	if len(os.Args) == 3 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		res := 0
// 		for _, a := range s1 {
// 			for _, b := range s2 {
// 				if a == b {
// 					isUniq := true
// 				}
// 				for _, c := range res {
// 					if a == c {
// 						isUniq = false
// 					}
// 				}
// 				if isUniq {

// 				}
// 			}
// 		}
// 	}
// }

// if len(os.Args) == 3 {
// 		s1 := os.Args[1]
// 		s2 := os.Args[2]
// 		res := ""
// 		for _, el1 := range s1 {
// 			for _, el2 := range s2 {
// 				if el1 == el2 {
// 					isUniq := true
// 					for _, el3 := range res {
// 						if el3 == el1 {
// 							isUniq = false
// 						}
// 					}
// 					if isUniq {
// 						res += string(el1)
// 					}
// 				}
// 			}
// 		}
// 		os.Stdout.WriteString(res + "\n")
// 	}
//*******************************
// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 2 {
// 		arr := []rune{}
// 		for _, x := range args[0] {
// 			for y := 0; y < len(args[1]); y++ {
// 				if byte(x) == args[1][y] && isIn(x, arr) {
// 					arr = append(arr, x)
// 				}
// 			}
// 		}
// 		for _, x := range arr {
// 			z01.PrintRune(x)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('\n')
// 	}
// 	// return
// }

// func isIn(char rune, arr []rune) bool {
// 	for _, x := range arr {
// 		if x == char {
// 			return false
// 		}
// 	}
// 	return true
// }
