package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	// defer z01.PrintRune('\n')
// 	counter := 0
// 	if len(os.Args) != 3 {
// 		return
// 	}
// 	if os.Args[1] == os.Args[2] {
// 		z01.PrintRune('1')
// 		z01.PrintRune(10)
// 		return
// 	}
// 	for _, v := range os.Args[2] {
// 		if len(os.Args[1]) == counter {
// 			z01.PrintRune('1')
// 			z01.PrintRune(10)
// 			return
// 		}
// 		if v == rune(os.Args[1][counter]) {
// 			counter++
// 		}
// 	}
// 	z01.PrintRune('0')
// 	z01.PrintRune(10)
// }

func main() {
	if len(os.Args) == 3 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		i := 0 //like counter
		if arg1 == arg2 {
			z01.PrintRune('1')
			z01.PrintRune(10)
			return
		}
		for j := range arg2 {
			if arg1[i] == arg2[j] {
				i++
				if i == len(arg1) {
					z01.PrintRune('1')
					z01.PrintRune(10)
					return
				}
			}
		}
		z01.PrintRune('0')
		z01.PrintRune(10)
	}
}

//******************************************************2 var
// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return
// 	}
// 	count := 0
// 	i := 0
// 	for _, el := range arg[1] {
// 		if rune(arg[0][i]) == el {
// 			i++
// 			if i == len(arg[0]) {
// 				i = 0
// 				count++
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }
//****************************
// counter := 0
// 	if len(os.Args) != 3 {
// 		returngithub.com/01-edu/go-tests/lib/challenge":= range os.Args[2] {
// 		if len(os.Args[1]) == counter {
// 			z01.PrintRune('1')
// 			z01.PrintRune(10)
// 			return
// 		}
// 		if v == rune(os.Args[1][counter]) {
// 			counter++
// 		}
// 	}
// 	z01.PrintRune('0')
// 	z01.PrintRune(10)
