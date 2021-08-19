package piscine

func Lcm(first, second int) int {
	mcm := first * second      // min common multiple
	min := mcm                 // temperory
	for i := mcm; i > 0; i-- { //цикл на уменьшение
		if i%first == 0 && i%second == 0 {
			if i < min {
				min = i
			}
		}
	}
	return min
}
