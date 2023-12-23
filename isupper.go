package piscine

func IsUpper(s string) bool {
	var index int
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			index++
		}
	}
	if len(s) == index {
		return true
	}
	return false
}
