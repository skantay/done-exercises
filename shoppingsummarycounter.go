package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for _, v := range str {
		if v != ' ' && v != '0' {
			word += string(v)
		} else {
			result[word]++
			word = ""
		}
	}
	result[word]++
	return result
}
