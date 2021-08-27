package main

import (
	"fmt"
)

func begin2() {
	var a int // сторона квадрата
	var S int //площадь квадрата
	fmt.Print("Введите сторону квадрата (a): ")
	fmt.Scan(&a)
	S = a * a

	fmt.Println(S)

}
