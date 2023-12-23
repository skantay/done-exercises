package piscine

func Capitalize(s string) string {
	words := []rune(s)
	capitalizeNext := true

	for i := 0; i < len(words); i++ {
		if isAlphanumeric(words[i]) {
			if capitalizeNext {
				words[i] = toUpperCase(words[i])
				capitalizeNext = false
			} else {
				words[i] = toLowerCase(words[i])
			}
		} else {
			capitalizeNext = true
		}
	}

	return string(words)
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toUpperCase(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 32
	}
	return char
}

func toLowerCase(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	}
	return char
}
