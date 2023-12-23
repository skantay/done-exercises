package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}

	for i := nb; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
