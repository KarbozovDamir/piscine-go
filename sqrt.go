package piscine

func Sqrt(nb int) int { //nb=16
	for i := 0; i <= nb; i++ { // i=0 i <=16
		if i*i == nb { // 0*0 != 16 and then there is an iteration, that is, already i = 1, 0 * 1! = 16 is performed again, and so on as i increases
			return i
		}
	}
	return 0
}
