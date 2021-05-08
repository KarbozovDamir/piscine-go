package piscine

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var arr []int
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}

	return arr
}
