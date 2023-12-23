package piscine

func Atoi(s string) int {
	var result int
	sign := 1

	for i, v := range s {
		if i == 0 && (v == '-' || v == '+') {
			if v == '-' {
				sign = -1
			}
			continue
		}
		if v < '0' || v > '9' {
			return 0
		}
		digit := int(v - '0')

		result += digit

		if i != len(s)-1 {
			result = result * 10
		}
	}

	return result * sign
}
