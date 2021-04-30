package piscine

func StrRev(s string) string {
	var ans string
	for _, c := range s {
		ans = string(c) + ans
	}
	return ans
}
