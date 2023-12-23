package piscine

func BasicAtoi2(s string) int {
	var result int

	for i, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		digit := int(v - '0')

		result += digit

		if i != len(s)-1 {
			result = result * 10
		}
	}

	return result
}
