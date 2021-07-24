package piscine

import "github.com/01-edu/z01"

// PrintComb2 : prcomb2

func PrintComb2() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i < j {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))

				z01.PrintRune(rune(j/10 + 48))
				z01.PrintRune(rune(j%10 + 48))
				if i != 98 || j != 99 {
					z01.PrintRune(',')
					z01.PrintRune(' ')

				}

			}
		}
	}
	z01.PrintRune('\n')
}

// ++
// +
// +
// ++
// +
// +
// +

// ++
// +
// +
// +
// +
// +
// +
// +
// +
// +
// ++

// /+/ ++
//+ +
// ++
// +
// ++
// ++
// +
// +
// +
// +

// ++
// +
// +

// +
// +
// ++
// +
// +

// ++
// +
// +
// +++
// +
// +
// +
// +
// ++
// +
// +
// ++++++++
// +++++++++++++++++++++++++++++++++++

// func PrintComb2() {
// 	for i := 0; i <= 99; i++ {
// 		for j := 0; j <= 99; j++ {
// 			if i < j {
// 				z01.PrintRune(rune(i/10 + 48))
// 				z01.PrintRune(rune(i%10 + 48))
// 				z01.PrintRune(' ')
// 				z01.PrintRune(rune(j/10 + 48))
// 				z01.PrintRune(rune(j%10 + 48))
// 				if i != 98 || j != 99 {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')
// 				}
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func PrintComb2() {
// 	for x := 0; x <= 99; x++ {
// 		for y := 0; y <= 99; y++ {
// 			if x < y {
// 				printR(inttorune(x))
// 				z01.PrintRune(' ')
// 				printR(inttorune(y))
// 				if x == 98 && y == 99 {
// 					z01.PrintRune('\n')
// 				} else {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')

// 				}
// 			}
// 		}
// 	}
// }

// func inttorune(num int) []rune {
// 	ans := []rune{}
// 	ans = append(ans, rune(num/10+48))
// 	ans = append(ans, rune(num%10+48))
// 	return ans
// }

// func printR(s []rune) {
// 	for _, i := range s {
// 		z01.PrintRune(i)
// 	}
// }

// 	for i := 0; i <= 99; i++ {
// 		for j := 0; j <= 99; j++ {
// 			if i < j {
// 				firstDigitOfI := i / 10
// 				secondDigitOfI := i % 10

// 				firstDigitOfJ := j / 10
// 				secondDigitOfJ := j % 10

// 				z01.PrintRune(rune(firstDigitOfI + '0'))
// 				z01.PrintRune(rune(secondDigitOfI + '0'))
// 				z01.PrintRune(rune(' '))
// 				z01.PrintRune(rune(firstDigitOfJ + '0'))
// 				z01.PrintRune(rune(secondDigitOfJ + '0'))
// 				if i == 98 && j == 99 {
// 					z01.PrintRune('\n')
// 				} else {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')
// 				}
// 			}
// 		}
// 	}
// }

// 	z01.PrintRune(10)
// }
// 	_min := '0'
// 	_max := '9'
// 	var st int = 0
// 	for i := _min; i <= _max; i++ {
// 		for j := _min; j <= _max; j++ {
// 			for k := i; k <= _max; k++ {
// 				for l := _min; l <= _max; l++ {
// 					if i == k && (j > l || j == l) {
// 						continue
// 					}
// 					if st > 0 {
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}
// 					z01.PrintRune(i)
// 					z01.PrintRune(j)
// 					z01.PrintRune(' ')
// 					z01.PrintRune(k)
// 					z01.PrintRune(l)
// 					st = 1
// 				}
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
