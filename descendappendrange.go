package piscine

func DescendAppendRange(max, min int) []int {
	if max < min || (max == 0 && min == 0) {
		return []int{}
	}
	result := []int{}

	for i := 0; max > min; i++ {
		result = append(result, max)
		max--
	}

	return result
}
