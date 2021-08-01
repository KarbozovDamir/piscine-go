package piscine

// FindNextPrime : prime
func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	for idx := 2; idx*idx <= nb; idx++ {
		if nb%idx == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
