package main

import (
	"fmt"
)

// func main() {
// 	s := "HelloHAHEhowHAHEareHAHEyou?"
// 	fmt.Printf("%#v\n", Split(s, "HAHE"))
// }
// func Split(s, sep string) []string {
// 	pool := s
// 	temp := ""
// 	slice := []string{}
// 	for i := 0; i < len(pool); i++ {
// 		step := i + len(sep)
// 		if i+len(sep) >= len(pool) {
// 			step = len(pool) - 1
// 		}
// 		temp = pool[i:step]
// 		fmt.Println(temp)
// 		if temp == sep {

// 			slice = append(slice, pool[0:step-len(sep)])
// 			pool := pool[step-len(sep):]
// 		}

// 	}
// 	return slice
// }

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

// firstHAsecondHAthridHA
func Split(s, sep string) []string {
	prev := 0 // H_..o
	ans := make([]string, 0)
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			ans = append(ans, s[prev:i])
			i += len(sep) // for skip HA
			prev = i      // FOR START AFTER SKIP HA
		}
	}
	if s[prev:] != sep {
		ans = append(ans, s[prev:])
	}
	return ans
}
