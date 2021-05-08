package piscine

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var res []int
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
