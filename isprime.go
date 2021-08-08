package piscine

// func IsPrime(nb int) bool {
// 	if nb < 0 {
// 		return false
// 	}
// 	if nb <= 3 {
// 		return nb > 1
// 	}
// 	if nb%2 == 0 || nb%3 == 0 {
// 		return false
// 	}
// 	i := 5
// 	for i*i <= nb {
// 		if nb%i == 0 || nb%(i+2) == 0 {
// 			return false
// 		}
// 		i = i + 6
// 	}
// 	return true
// }

// func FindNextPrime(nb int) int {
// 	if nb < 3 {
// 		return 2
// 	}
// 	for id := 2; id*id <= nb; id++ {
// 		if nb%id == 0 {
// 			return FindNextPrime(nb + 1)
// 		}
// 	}
// 	return nb
// }

// IsPrime : prime
func IsPrime(nb int) bool {
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
