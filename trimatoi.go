package piscine

func TrimAtoi(s string) int {
	var arr string
	for _, item := range s {
		if item == '-' && Len(arr) == 0 {
			arr = string(item) + arr
		}
		if item >= 48 && item <= 57 {
			arr = arr + string(item)
		}
	}
	return Atoi(arr)
}
