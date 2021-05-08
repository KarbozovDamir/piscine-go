package piscine

func MakeRange(min, max int) []int {
	if max-min > 0 {
		res := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			res[i] = i + min
		}
		return res
	} else {
		return nil
	}
}
