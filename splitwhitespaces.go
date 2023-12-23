package piscine

func SplitWhiteSpaces(s string) []string {
	words := []rune(s)
	result := []string{}
	var word string
	for i := 0; i < len(s); i++ {
		if words[i] == ' ' || i == len(s)-1 {
			if i == len(s)-1 {
				word += string(words[i])
			} else if words[i+1] == ' ' {
				continue
			}
			result = append(result, word)
			word = ""
			continue
		}
		word += string(words[i])
	}
	return result
}
