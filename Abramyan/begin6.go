package main

import (
	"fmt"
)

func begin6() {
	var a, b, c, V, S int
	// fmt.Scan(&a, &b, &c)
	fmt.Scan(&a)
	fmt.Println("Ребро a: ", a)
	fmt.Scan(&b)
	fmt.Println("Ребро b: ", b)
	fmt.Scan(&c)
	fmt.Println("Ребро c: ", c)
	V = a * b * c
	fmt.Println("Объем: ", V)
	S = 2 * (a*b + b*c + a*c)
	fmt.Println("Площадь поверхности: ", S)
}

// Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти
// его объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).
