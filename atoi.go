package piscine

// Atoi function converts string to int
func Atoi(s string) int {
	numbers := []rune(s)
	var result int
	var isNegative bool

	if numbers[0] == '-' {
		numbers = numbers[1:]
		isNegative = true
	}

	for i := 0; i < len(numbers); i++ {

		if numbers[i] >= '0' && numbers[i] <= '9' {
			result = (result * 10) + int(numbers[i] - '0')			
			continue
		}
		return 0
	}
	if isNegative {
		result = -result
	}
	return result
}