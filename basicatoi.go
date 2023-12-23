package piscine

func BasicAtoi(s string) int {
	var result int

	for i, v := range s {

		digit := int(v - '0')

		result += digit

		if i != len(s)-1 {
			result = result * 10
		}
	}

	return result
}
