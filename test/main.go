package main

import "fmt"

func main() {
	fmt.Println(atoi("582"))
	fmt.Println(itoa(582))
}

func atoi(s string) int {
	sum := 0
	for _, num := range s {
		sum = int(num-48) + sum*10
	}
	return sum
}
func itoa(n int) string {
	str := ""
	for n != 0 {
		ostatok := n % 10
		str = string(ostatok+48) + str
		n = n / 10
	}
	return str

}
