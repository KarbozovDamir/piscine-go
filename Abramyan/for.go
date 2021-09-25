package main

import "fmt"

func main() {
	for40()
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

func for40() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	sum := 1               // A должно выводиться 1 раз
	for ; n1 <= n2; n1++ { // Вывести все целые числа от A до B включительно
		for i := 0; i < sum; i++ { //при этом число A должно выводиться 1 раз, число A + 1 должно выводиться 2 раза и т. д.
			fmt.Println(n1)
		}
		sum++
	}
}
