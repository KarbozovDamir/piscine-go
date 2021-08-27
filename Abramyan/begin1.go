package main

import (
	"fmt"
)

func begin1() {
	var a int
	var p int //Периметр квадрата
	fmt.Print("Введите сторону квадрата (a): ")
	fmt.Scan(&a)
	p = 4 * a

	fmt.Println(p)

}
