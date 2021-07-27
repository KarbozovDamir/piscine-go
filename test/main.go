package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))

	// s := "Hello World!"
	// s = piscine.StrRev(s)
	// fmt.Println(s)

	// piscine.PrintCombN(1)
	// piscine.PrintCombN(3)
	// piscine.PrintCombN(9)

	// s := "Hello World!"
	// s = piscine.StrRev(s)
	// fmt.Println(s)

}

// 	x := []int{}
// 	min := x[0]
// 	for _, el := range x {
// 		if el < min {
// 			min = el
// 		}
// 	}
// 	fmt.Println(min)
// }

//
// func main() {
// 	for i := 0; i < 5; i++ {
// 		for j := 5 - i; j > 0; j-- {
// 			fmt.Print("*")
// 		}
// }
// }
