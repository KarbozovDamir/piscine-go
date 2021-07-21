package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				firstDigitOfI := i / 10
				secondDigitOfI := i % 10

				firstDigitOfJ := j / 10
				secondDigitOfJ := j % 10

				z01.PrintRune(rune(firstDigitOfI + '0'))
				z01.PrintRune(rune(secondDigitOfI + '0'))
				z01.PrintRune(rune(' '))
				z01.PrintRune(rune(firstDigitOfJ + '0'))
				z01.PrintRune(rune(secondDigitOfJ + '0'))
				if i == 98 && j == 99 {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}

		}
	}
}

// 	for i := '0'; i <= '9'; i++ {
// 		for j := '0'; j <= '9'; j++ {
// 			for k := i; k <= '9'; k++ {
// 				for l := '1'; l <= '9'; l++ {
// 					if i == k && l <= j {
// 						continue
// 					}
// 					z01.PrintRune(i)
// 					z01.PrintRune(j)
// 					z01.PrintRune(' ')
// 					z01.PrintRune(k)
// 					z01.PrintRune(l)
// 					if i != '9' || j != '8' || k != '9' || l != '9' {
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 		}
// 	}
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
