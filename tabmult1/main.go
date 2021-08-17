package main

import (
	"os"
)

func atoi(s string) int {
	var res int = 0
	for _, el := range s {
		if el >= 48 && el <= 57 {
			res = res*10 + int(el-48) // +9...807 max
		} else {
			return 0
		}
	}
	return res
}

func itoa(n int) string {
	str := ""
	for n != 0 {
		str = string(rune(n%10)+48) + str
		n /= 10
	}
	return str
}

func main() {
	args := os.Args[1:]
	//sum := ""
	length := 0
	for range args {
		length++
	}
	if length != 1 {
		return
	}
	num := atoi(args[0])
	for i := 1; i <= 9; i++ {
		os.Stdout.WriteString(itoa(i) + " x " + itoa(num) + " = " + itoa(i*num) + "\n")
		// 	sum += itoa(i)
		// 	sum += " x "
		// 	sum += itoa(num)
		// 	sum += " = "
		// 	sum += itoa(i * num)
		// 	sum += "\n"
		// }
		// for _, h := range sum {
		// 	z01.PrintRune(h)
		// }
	}

}
