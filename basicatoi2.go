package piscine

// Atoi function converts string to int
func BasicAtoi2(s string) int {
	numbers := []rune(s)
	var result int

	for i := 0; i < len(numbers); i++ {

		if numbers[i] >= '0' && numbers[i] <= '9' {
			result = (result * 10) + int(numbers[i] - '0')			
			continue
		}
		return 0
	}
	return result
}