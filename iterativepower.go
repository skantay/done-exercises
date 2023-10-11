package piscine

func IterativePower(nb int, power int) int {
	var result int = 1
	for 0 < power {
		result *= nb
		power--
	}
	return result
}