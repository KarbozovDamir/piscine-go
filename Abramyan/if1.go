package main

//Даны три числа. Найти среднее из них (то есть число, расположенное
// между наименьшим и наибольшим).(13)
// func main() {
// 	var a, b, c int
// 	fmt.Scan(&a)
// 	fmt.Println("a" + "\n")
// 	fmt.Scan(&b)
// 	fmt.Println("b" + "\n")
// 	fmt.Scan(&c)
// 	fmt.Println("c" + "\n")
// 	if a < b && a < c || a < 0 {
// 		fmt.Println("a наименьшее")
// 		return
// 	}
//**********************************************************************************************
// Даны три числа. Найти наименьшее из них.(12)
// func main() {
// 	var a, b, c, mn int
// 	fmt.Scan(&a)
// 	fmt.Println("a" + "\n")
// 	fmt.Scan(&b)
// 	fmt.Println("b" + "\n")
// 	fmt.Scan(&c)
// 	fmt.Println("c" + "\n")
// 	if a < b {
// 		mn = a
// 		fmt.Println("Минимум:", mn)
// 	} else {
// 		mn = b
// 		fmt.Println("Минимум:", mn)
// 	}
// 	if mn > c {
// 		mn = c
// 		fmt.Println("Минимум:", mn)
// 	}
// 	return
// }

//Даны три целых числа. Найти количество положительных и количество
// отрицательных чисел в исходном наборе.
// func main() {
// 	var a, b, c, sumPlus, sumMinus int
// 	fmt.Scan(&a, &b, &c)
// 	if a > 0 || b > 0 || c > 0 {
// 		sumPlus = a + b + c
// 		fmt.Println(sumPlus)
// 	}
// 	if a < 0 || b < 0 || c < 0 {
// 		sumMinus = a + b + c
// 		fmt.Println(sumMinus)
// 	}

// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	if n > 0 {
// 		n++
// 		fmt.Println(n)
// 	}
// 	if n < 0 {
// 		n -= 2
// 		fmt.Println(n)
// 	}
// 	if n == 0 { // если является нулевым
// 		n = 10 // тогда замени на 10
// 		fmt.Println(n)

// 	}
// }

// Дано целое число. Если оно является положительным, то прибавить к
// нему 1; если отрицательным, то вычесть из него 2; если нулевым, то
// заменить его на 10. Вывести полученное число.

//////////////////////////////////////////////////////////////////////////////////////////
/*Дано целое число. Если оно является положительным, то прибавить к
нему 1; в противном случае вычесть из него 2. Вывести полученное число.*/
// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	if n > 0 {
// 		n++
// 		fmt.Println(n)
// 	} else {
// 		n -= 2
// 		fmt.Println(n)
// 	}
// }
/////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	var a int
// 	fmt.Scan(&a)
// 	if a > 0 {
// 		a++
// 	}
// 	fmt.Println(a)

// }
// }
// Дано целое число. Если оно является положительным, то прибавить к
// нему 1; в противном случае не изменять его. Вывести полученное число
