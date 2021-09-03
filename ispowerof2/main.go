// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func isPowerOf2(n int) bool {
// 	return n != 0 && ((n & (n - 1)) == 0)
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr, _ := strconv.Atoi(os.Args[1])
// 		fmt.Println(isPowerOf2(nbr))
// 	}
// }

package main

import (
	"os"
)

func atoi(s string) (num int) {
	// num := 0
	for i := 0; i < len(s); i++ {
		num = int(s[i]-48) + num
		if i != len(s)-1 {
			num *= 10
		}
	}
	return num
}

func main() {
	if len(os.Args) == 2 {
		a := atoi(os.Args[1])
		if a%2 != 0 {
			os.Stdout.Write([]byte("false\n"))
		} else {
			for i := 2; i < a; i *= 2 {
				if 2*i == a {
					os.Stdout.Write([]byte("true\n"))
					return
				}
			}
			os.Stdout.Write([]byte("false\n"))
		}
	}
}
