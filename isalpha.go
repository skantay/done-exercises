package piscine

func IsAlpha(s string) bool {
	var count int
	for _, v := range s {
		if ((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) || (v >= '0' && v <= '9') {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
