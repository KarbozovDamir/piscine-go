package piscine

// func main() {
// 	fmt.Println(Nauuo(50, 43, 20))
// 	fmt.Println(Nauuo(13, 13, 0))
// 	fmt.Println(Nauuo(10, 9, 0))
// 	fmt.Println(Nauuo(5, 9, 2))
// }
func Nauuo(plus, minus, rand int) string {
	if plus > minus+rand {
		return "+"
	} else if minus > plus+rand {
		return "-"
	} else if plus == minus && rand == 0 {
		return "0"
	} else {
		return "?"
	}
}

/************************************************2 variant
func Nauuo(plus, minus, rand int) string {
	res := ""
	diferent := plus - minus
	if diferent < 0 {
		diferent = minus - plus
	}
	if diferent == 0 && rand == 0 {
		return "0"
	}
	if diferent <= rand {
		return "?"
	} else if plus == minus {
		return "0"
	} else if plus > minus {
		return "+"
	} else if plus < minus {
		return "-"
	}
	return res
}*/
