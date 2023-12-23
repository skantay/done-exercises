package piscine

func TrimAtoi(s string) int {
	var result int
	var isnegative bool

	for _, v := range s {
		if v == '-' && result == 0 {
			isnegative = true
		}
		if v >= '0' && v <= '9' {
			result *= 10
			num := int(v - '0')
			result = result + num

		}
	}
	if isnegative {
		return result * -1
	}
	return result
}
