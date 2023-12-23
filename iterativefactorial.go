package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}
	var result int = 1
	for nb != 0 {
		result *= nb

		nb--
	}

	return result
}
