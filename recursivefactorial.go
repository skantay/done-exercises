package piscine

func RecursiveFactorial(nb int) int {

	if nb == 0 { // If nb == 0 return 1
		return 1
	}
	
	return nb * RecursiveFactorial(nb-1) // nb will be multiplied by nb - 1 each time
}

// It seems to be overflowed after nb 21
// factorial of 1 or 0 = 1