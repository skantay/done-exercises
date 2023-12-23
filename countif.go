package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, word := range tab {
		if f(word) {
			count++
		}
	}
	return count
}
