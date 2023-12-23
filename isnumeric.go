package piscine

func IsNumeric(s string) bool {
	var count int
	for _, v := range s {
		if v >= '0' && v <= '9' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
