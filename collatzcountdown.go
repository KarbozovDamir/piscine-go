package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	steps := 2

	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else if {
			if start%2 == 1
			start = 3*start + 1
		}
		steps++
	}
	return steps
}
