package piscine

func IsLower(s string) bool {
	var index int
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			index++
		}
	}
	if len(s) == index {
		return true
	}
	return false
}
