package piscine

func TrimAtoi(s string) int {
	res1 := 0
	res2 := 1
	for _, item := range s {
		if item >= 48 && item <= 57 {
			res1 = (int(item) - 48) + res1*10
		} else if item == '-' && res1 == 0 {
			res2 = -1
		}
	}
	return res1 * res2
}
