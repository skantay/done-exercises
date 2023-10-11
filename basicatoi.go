package piscine

// Atoi function converts string to int
func BasicAtoi(s string) int {
	numbers := []rune(s)
	var result int
	for i := 0; i < len(numbers); i++ {
		result = (result * 10) + int(numbers[i] - '0')
	}
	return result
}