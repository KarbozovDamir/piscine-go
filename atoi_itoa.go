package piscine

// func main() {
// 	fmt.Println(atoi("582"))
// 	fmt.Println(itoa(582))
// }

func atoi(s string) int {
	sum := 0                // создаем пустой результат int
	for _, num := range s { //разбиваем s на num
		n := int(num - 48) // получаем число из ASCII ex: 53-48 = 5
		sum = n + sum*10   // сумму увеличиваем на разряд и плюсуем число
	}
	return sum // возвращаем результат (int)
}
func itoa(n int) string {
	str := "" // создаем пустой результат string
	for n != 0 {
		ostatok := n % 10           // находим последнюю цифру
		ostatok = ostatok + 48      // остаток переводми в ASCII ex: 5+48 = 53
		str = string(ostatok) + str // перевожу в string и плюсую к остатку 48, затем прибавляю к str
		n = n / 10                  // уменьшаю на разряд
	}
	return str // возвращаем результат (string)

}
