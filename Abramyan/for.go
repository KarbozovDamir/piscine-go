package main

import "fmt"

func main() {
	arr1()
}

//**********1 esep
func for1() {

	var k, n int
	fmt.Scan(&k)
	fmt.Scan(&n) // 5

	for ; 0 < n; n-- { // 5 4  3 2 1 0
		fmt.Printf("%d\n", k)
	}
	// for i := 1; i <= n; i++ {
	// 	fmt.Printf("%d\n", k)
	// }
}

func for2() {
	var A, B int
	fmt.Scan(&A, &B)
	// fmt.Scan(&n)
	for ; A <= B; A++ {
		fmt.Println(A)
	}
}

/*Даны целые числа A и B (A < B). Вывести все целые числа от A до B включительно; при этом число A должно выводиться 1 раз, число A + 1 должно выводиться 2 раза и т. д.*/
func for40() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	sum := 1               // A должно выводиться 1 раз
	for ; n1 <= n2; n1++ { // Вывести все целые числа от A до B включительно т.е n2 -n1 + 1 {считаем количество чисел}
		for i := 0; i < sum; i++ { //при этом число A должно выводиться 1 раз, число A + 1 должно выводиться 2 раза и т. д.
			fmt.Println(n1)
		}
		sum++
	}
}

/*
Даны целые положительные числа A и B (A < B). Вывести все целые числа от A до B включительно;
при этом каждое число должно выводиться столько раз, каково его значение (например, число 3 выводится 3 раза).*/

func for39() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		for j := 0; j < i; j++ {

			fmt.Print(i)
		} 
		fmt.Println()
	}
}

/*Дано целое число N (> 0). Найти сумму 1N + 2N–1 + … + N1. Чтобы избежать целочисленного переполнения, вычислять слагаемые этой суммы с помощью вещественной
переменной и выводить результат как вещественное число.*/
// func for38() {

// }
func arr1() {
	var n int{
	// c = 2

	fmt.Print("N:")
	fmt.Scan(&n)

	a := make([]int, n)

	for i := len(a) - 1; i < n; i++ {
		// i++
		//c++
		fmt.Print(a)
	}

	// fmt.Println()

}
func for00(){
	
} 
// var n int
// fmt.Scan(&n)
// arr := []int{}
// for i := len(arr) - 1; i >= n; i-- {

// 	fmt.Println(arr[i])

// }
// }
