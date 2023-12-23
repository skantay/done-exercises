package piscine

func Any(f func(string) bool, a []string) bool {
	for _, word := range a {
		if f(word) {
			return true
		}
	}
	return false
}
