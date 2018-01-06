package intToRoman

func intToRoman(num int) string {
	roman := ""
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strings := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, v := range values {
		for num >= v {
			num -= values[i]
			roman += strings[i]
		}
	}

	return roman
}
