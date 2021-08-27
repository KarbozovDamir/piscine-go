package main

import (
	"fmt"
)

func begin3() {
	var a, b int //стороны прямоугольника
	var S int    // площадь прямоугольника
	var P int    // Периметр прямоугольника
	fmt.Print("Введите сторону a: ")
	fmt.Scan(&a)
	fmt.Print("Введите сторону b: ")
	fmt.Scan(&b)
	S = a * b
	P = 2 * (a + b)
	fmt.Println("Площадь прямоугольника S = ", S)
	fmt.Println("Периметр прямоугольника P = ", P)
}
