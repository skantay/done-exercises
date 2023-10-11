package piscine

func IterativeFactorial(nb int) int {
	var result int = 1

	for ;nb != 1;nb--{
		result *= nb
	}
	return result
}

// It seems that it get overflowed starting from 21