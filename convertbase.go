package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// atoi

	var result string
	var buffer int

	// decimal to baseto

	buffer = convertToDecimal(nbr, baseFrom)

	if baseTo == "0123456789" {
		return itoa(buffer)
	}

	result = convertToBase(buffer, baseTo)

	return result
}

func convertToBase(buffer int, baseTo string) string {
	var result string
	baseLen := len(baseTo)

	for buffer != 0 {
		digit := getBaseDigit((buffer % baseLen), baseTo)
		result = digit + result
		buffer /= baseLen
	}

	return result
}

func getBaseDigit(index int, baseTo string) string {
	return string(baseTo[index])
}

func convertToDecimal(nbr, baseFrom string) int {
	baseLen := len(baseFrom)
	sLen := len(nbr)

	var result int

	for _, v := range nbr {

		digit, ok := getDigit(v, baseFrom)

		if !ok {
			return 0
		}

		poweredNum := pwr(baseLen, sLen)
		result += (digit * poweredNum)

		sLen--
	}

	return result
}

func getDigit(s rune, baseFrom string) (int, bool) {
	baseRunes := []rune(baseFrom)

	for i := 0; i < len(baseRunes); i++ {
		if s == baseRunes[i] {
			return i, true
		}
	}

	return 0, false
}

func pwr(n, power int) int {
	if power == 1 {
		return power
	} else {
		return n * pwr(n, power-1)
	}
}

func atoi(n string) int {
	var result int

	for i, v := range n {
		result += int(v - '0')
		if i != len(n)-1 {
			result *= 10
		}
	}

	return result
}

func itoa(n int) string {
	var result string

	for n != 0 {
		result = string(rune(n%10)+'0') + result
		n /= 10
	}

	return result
}
