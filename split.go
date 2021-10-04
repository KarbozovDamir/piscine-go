package piscine

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

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
	if prev != len(s)-1 {
		ans = append(ans, s[prev:])
	}
	return ans
}
