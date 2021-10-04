package piscine

// func Lcm(first, second int) int {
// 	mcm := first * second      // min common multiple
// 	min := mcm                 // temperory
// 	for i := mcm; i > 0; i-- { //цикл на уменьшение
// 		if i%first == 0 && i%second == 0 {
// 			if i < min {
// 				min = i
// 			}
// 		}
// 	}
// 	return min
// }
func Lcm(first, second int) int {
	res := 0
	for i := second; i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			break
		}
	}
	return res
}
