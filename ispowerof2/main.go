package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	if len(os.Args) == 2 {
		n, _ := strconv.Atoi(os.Args[1])

		if n == 1 || n == 2 {
			out("true")
			return
		}

		for n > 1 {
			if n%2 == 1 {
				out("false")
				return
			}
			n /= 2
		}
		out("true")
	}
}
func out(s string) {
	for _, el := range s {
		z01.PrintRune(el)
	}
	z01.PrintRune(10)
}

// isPowerOf2
// func main() {
// 	if len(os.Args[1:]) != 1 {
// 		return
// 	}
// 	x, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		return
// 	}

// 	for x > 1 {
// 		if x%2 == 1 {
// 			fmt.Println("false")
// 			return
// 		}
// 		x /= 2
// 	}
// 	fmt.Println("true")
// }

//************************1 var  with bit operation
// func isPowerOf2(n int) bool {
// 	return n != 0 && ((n & (n - 1)) == 0)
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr, _ := strconv.Atoi(os.Args[1])
// 		fmt.Println(isPowerOf2(nbr))
// 	}
// }

//***************************************************var 1
// func atoi(s string) (num int) {
// 	// num := 0
// 	for i := 0; i < len(s); i++ {
// 		num = int(s[i]-48) + num
// 		if i != len(s)-1 {
// 			num *= 10
// 		}
// 	}
// 	return num
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	if len(os.Args) == 2 {
// 		a := atoi(os.Args[1])
// 		if a%2 != 0 {
// 			os.Stdout.Write([]byte("false\n"))
// 		} else {
// 			for i := 2; i < a; i *= 2 {
// 				if 2*i == a {
// 					os.Stdout.Write([]byte("true\n"))
// 					return
// 				}
// 			}
// 			os.Stdout.Write([]byte("false\n"))
// 		}
// 	}
// }
