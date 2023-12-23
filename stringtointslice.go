package piscine

func StringToIntSlice(str string) []int {
	var result []int

	for _, v := range str {
		result = append(result, int(v))
	}

	return result
}
