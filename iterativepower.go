package piscine

func IterativePower(nb int, power int) int {
	p := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := power; i > 1; i-- {
		p = nb * p
	}
	return p
}