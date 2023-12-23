package piscine

func AppendRange(min, max int) []int {
	if max < min || (max == 0 && min == 0) {
		return nil
	}
	result := []int{}

	for i := 0; min < max; i++ {
		result = append(result, min)
		min++
	}

	return result
}
