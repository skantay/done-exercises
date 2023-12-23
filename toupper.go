package piscine

func ToUpper(s string) string {
	var result string
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v -= 32 // v = v - 32
		}
		result = result + string(v)
	}
	return result
}
