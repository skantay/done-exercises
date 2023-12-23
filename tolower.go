package piscine

func ToLower(s string) string {
	var result string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v += 32
		}
		result = result + string(v)
	}
	return result
}
