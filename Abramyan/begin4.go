package main

import (
	"fmt"
)

const pi = 3.14

func begin() {
	var d float64 //диаметр окружности
	fmt.Scanln(&d)
	L := pi * d //найти длину окружности
	fmt.Println(L)
}
