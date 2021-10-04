package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, r := range "Error" {
			z01.PrintRune(r)
		}
		z01.PrintRune(10)
	}
	if len(os.Args) == 3 {
		num1, _ := strconv.Atoi(os.Args[1])
		num2, _ := strconv.Atoi(os.Args[2])
		if num2 == 0 && num1 == 0 {
			for _, r := range "Error" {
				z01.PrintRune(r)
			}
			z01.PrintRune(10)
		}
		for i := 0; i < num2; i++ {
			for j := 0; j < num1; j++ {
				if i%2 == 0 {
					if j%2 == 0 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
				} else {
					if j%2 == 0 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('#')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}

// func main() {
// 	if len(os.Args) != 3 {
// 		for _, r := range "Error" {
// 			z01.PrintRune(r)
// 		}
// 		z01.PrintRune(10)
// 		return
// 	}
// 	if len(os.Args) == 3 {
// 		n1, _ := strconv.Atoi(os.Args[1])
// 		n2, _ := strconv.Atoi(os.Args[2])

// 		for i := 1; i <= n1; i++ {
// 			for j := 1; j <= n2; j++ {
// 				if i%2 == 1 && j%2 == 1 {
// 					z01.PrintRune('#')
// 				} else if i%2 == 1 && j%2 == 0 {
// 					z01.PrintRune(' ')
// 				} else if i%2 == 0 && j%2 == 1 {
// 					z01.PrintRune(' ')
// 				} else {
// 					z01.PrintRune('#')
// 				}
// 			}
// 		}
// 		// for _, r := range "Error" {
// 		// 	z01.PrintRune(r)
// 		// }
// 		// z01.PrintRune(10)
// 		// return
// 		z01.PrintRune('\n')
// 	}
// }

//******************************************var1
// func main() {
// 	if len(os.Args) != 3 {
// 		error()
// 		return
// 	}
// 	x, err := strconv.Atoi(os.Args[1])
// 	y, err2 := strconv.Atoi(os.Args[2])

// if err != nil || err2 != nil {
// 	error()
// 	return
// }
// 	if x == 0 || y == 0 {
// 		error()
// 		returnfor _, el := range flag {
// 	z01.PrintRune(el)
// }
// z01.PrintRune(10)
// return
// 		for j := 1; j <= x; j++ {
// 			if (i+j)%2 == 0 {
// 				z01.PrintRune('#')
// 			} else {
// 				z01.PrintRune(' ')
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func error() {
// 	for _, el := range "Error" {
// 		z01.PrintRune(el)
// 	}
// 	z01.PrintRune(10)
// }

//*************************************
// if len(os.Args) == 2 {
// 	for _, r := range "Error" {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune(10)
// }
//
// 	if len(os.Args) == 3 {
// 		num1, _ := strconv.Atoi(os.Args[1])
// 		num2, _ := strconv.Atoi(os.Args[2])
// 		if num2 == 0 && num1 == 0 {
// 			for _, r := range "Error" {
// 				z01.PrintRune(r)
// 			}
// 			z01.PrintRune(10)
// 		}
// 		for i := 0; i < num2; i++ {
// 			for j := 0; j < num1; j++ {
// 				if i%2 == 0 {
// 					if j%2 == 0 {
// 						z01.PrintRune('#')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				} else {
// 					if j%2 == 0 {
// 						z01.PrintRune(' ')
// 					} else {
// 						z01.PrintRune('#')
// 					}
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }
