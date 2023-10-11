package piscine

func IsPrime(nb int) bool {
	for i := 2; i * i <= nb; i++ {
		if nb % i == 0 { // if nb modulo i == zero
			return false
		}
	}
	return true
}