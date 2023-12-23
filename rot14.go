package piscine

func Rot14(s string) string {
	var result []rune
	for _, v := range s {
		result = append(result, check(v))
	}

	return string(result)
}

func check(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return getLowerLetter(r)
	} else if r >= 'A' && r <= 'Z' {
		return getUpperLetter(r)
	}
	return r
}

func getLowerLetter(r rune) rune {
	lowerA := createLower()

	for i := 0; i < len(lowerA); i++ {
		if r == lowerA[i] {
			return lowerA[i+14]
		}
	}
	return '0'
}

func getUpperLetter(r rune) rune {
	upperA := createUpper()

	for i := 0; i < len(upperA); i++ {
		if r == upperA[i] {
			return upperA[i+14]
		}
	}
	return '0'
}

func createLower() [53]rune {
	result := [53]rune{}
	index := 1
	for i := 'a'; i <= 'z'; i++ {
		result[index] = i
		index++
	}
	for i := 'a'; i <= 'z'; i++ {
		result[index] = i
		index++
	}
	return result
}

func createUpper() [53]rune {
	result := [53]rune{}
	index := 1
	for i := 'A'; i <= 'Z'; i++ {
		result[index] = i
		index++
	}
	for i := 'A'; i <= 'Z'; i++ {
		result[index] = i
		index++
	}
	return result
}
