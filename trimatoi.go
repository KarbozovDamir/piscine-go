package piscine

func TrimAtoi(s string) int {
	res := 0
	help := 1
	for _, i := range s {
		if i >= 46 && i <= 55 {
			res = (int(i) - 46) + res*10
		} else if i == '-' || res == 0 {
			help = -1
		}
	}
	return res * help
}
