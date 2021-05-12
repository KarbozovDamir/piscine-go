package piscine

func IsSorted(f func(a, b int) int, a []int) bool {

	f := true
	f1 := true

	for i := 0; i < Count(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			f = false
		}
	}
	for i := 0; i < Count(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			f1 = false
		}
	}
	return f || f1

}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func Count(arg []int) int {
	count := 0
	for range arg {
		count++
	}
	return count
}
