package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	if len(word) < n || n < 0 || n == 0 {
		return '\x00'
	}

	return word[n-1]
}
