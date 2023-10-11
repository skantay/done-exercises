package piscine

func RecursivePower(nb int, power int) int {
	if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1) // nb will be multiplied ny nb until power == 1 
	}
}