package piscine

import "github.com/01-edu/z01"

// PrintComb : () {
func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for k := '2'; k <= '9'; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i < '7' || j < '8' || k < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// 	for a := '0'; a <= '7'; a++ {
// 		for b := '1'; b <= '8'; b++ {
// 			for c := '2'; c <= '9'; c++ {
// 				if a < b && b < c {
// 					if a == '7' && b == '8' && c == '9' {
// 						z01.PrintRune(a)
// 						z01.PrintRune(b)
// 						z01.PrintRune(c)
// 					} else {
// 						z01.PrintRune(a)
// 						z01.PrintRune(b)
// 						z01.PrintRune(c)
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}
//

// 	for a := '0'; a <= '7'; a++ {
// 		for b := '1'; b <= '8'; b++ {
// 			for c := '2'; c <= '9'; c++ {
// 				if a < b && b < c {
// 					if a == '7' && b == '8' && c == '9' {
// 						z01.PrintRune(a)
// 						z01.PrintRune(b)
// 						z01.PrintRune(c)
// 					} else {
// 						z01.PrintRune(a)
// 						z01.PrintRune(b)
// 						z01.PrintRune(c)
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}var str int = 0
// for i := _min; i <= _max; i++ {
// 	for j := _min; j <= _max; j++ {
// 		for k := _min; k <= _max; k++ {
// 			if i < j && j < k {
// 				if str > 0 {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')
// 				}
// 				z01.PrintRune(i)
// 				z01.PrintRune(j)
// 				z01.PrintRune(k)
// 				str = 0
// 			}
// 		}
// 	}
// }
// }

// var str int = 0
// for i := _min; i <= _max; i++ {
// 	for j := _min; j <= _max; j++ {
// 		for k := _min; k <= _max; k++ {
// 			if i < j && j < k {
// 				if str > 0 {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')
// 				}
// 				z01.PrintRune(i)
// 				z01.PrintRune(j)
// 				z01.PrintRune(k)
// 				str = 0
// 			}
// 		}
// 	}
// }
// }
