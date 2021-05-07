package piscine

func Capitalize(s string) string {
	arr := []rune(s)
	count := 0
	for range arr {
		count++
	}
	if arr[0] >= 'a' && arr[0] <= 'z' {
		arr[0] -= 32
	}
	for i := 1; i < count; i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			if (arr[i-1] >= 'A' && arr[i-1] <= 'Z') || (arr[i-1] >= 'a' && arr[i-1] <= 'z') || (arr[i-1] >= '0' && arr[i-1] <= '9') {
				arr[i] += 32
			}
		}
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if (arr[i-1] >= 'A' && arr[i-1] <= 'Z') || (arr[i-1] >= 'a' && arr[i-1] <= 'z') || (arr[i-1] >= '0' && arr[i-1] <= '9') {
				continue
			} else {
				arr{i] -= 32
			}
		}
	}
	res := string(arr)
	return res
}
