package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, 0, len(a))

	for i := 0; i < len(a); i++ {
		result = append(result, f(a[i]))
	}

	return result
}
