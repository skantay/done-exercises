package piscine

func JumpOver(str string) string {
	words := []rune(str)
	var result string
	if len(words) < 2 {
		return "\n"
	}

	for i := 2; i < len(words); i = i + 3 {
		result += string(words[i])
	}
	result += "\n"
	return result
}
