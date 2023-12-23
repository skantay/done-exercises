package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	var index int = -1
	var found bool = false
	for _, val := range toFind {
		for i, v := range s {
			if val == v && !found {
				return i
			}
		}
		found = true
	}
	return index
}
