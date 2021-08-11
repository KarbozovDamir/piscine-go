package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
	fmt.Println(Atoi("12+34"))
	fmt.Println(Atoi("1234+"))
	fmt.Println(Atoi("-9223372036854775808"))
	fmt.Println(Atoi("9223372036854775807"))
	fmt.Println(Atoi("9223372036854775808"))
	fmt.Println(Atoi("9223372036854775807898616"))

}

func Atoi(s string) int {
	res := 0
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		neg = true
	}

	for _, el := range s {
		if el >= 48 && el <= 57 {
			res = res*10 + int(el-48)
		} else {
			return 0
		}
	}
	if neg {
		res *= -1
	}
	if !neg && res < 0 { // !neg -> false
		return 0
	}
	if neg && res > 0 { // true
		return 0
	}
	return res
}
