// Дана длина ребра куба a. Найти объем куба V = a
// 3 и площадь его
// поверхности S = 6·a
// 2
package main

import (
	"fmt"
)

func begin5() {
	var a, V, S int
	fmt.Scan(&a)
	V = a * a * a // Ob'yom kuba
	fmt.Println("Ob'yom kuba: V = ", V)
	S = 6 * a * a // Площадь поверхности
	fmt.Println("Площадь поверхности: S=", S)
}
