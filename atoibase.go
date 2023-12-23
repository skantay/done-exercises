package piscine

func AtoiBase(s string, base string) int {
	if notValidBase(base) {
		return 0
	}

	result := getValue(s, base)

	return result
}

func getValue(s, base string) int {
	baseLen := len(base)
	lenghtS := len(s)
	var result int

	for i := 0; i < len(s); i++ {
		digit, ok := getValueBase(rune(s[i]), base)

		if !ok {
			return 0
		}

		poweredNum := powerOf(baseLen, lenghtS)

		result = result + (digit * poweredNum)
		lenghtS--
	}

	return result
}

func powerOf(n, power int) int {
	if power == 1 {
		return power
	} else {
		return n * powerOf(n, power-1)
	}
}

func getValueBase(s rune, base string) (int, bool) {
	i := 0
	baseRunes := []rune(base)

	for ; i < len(baseRunes); i++ {
		if s == baseRunes[i] {
			return i, true
		}
	}

	return 0, false
}

// Check if base is valid
func notValidBase(base string) bool {
	if len(base) < 2 {
		return true
	}

	if isUnique(base) {
		return true
	}

	return false
}

// Check if base's chars are unique, and does not contain + - signs
func isUnique(base string) bool {
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				return true
			}
		}
	}
	return false
}
