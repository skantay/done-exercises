package piscine

func Compact(ptr *[]string) int {
	result := make([]string, 0, len(*ptr))
	for _, val := range *ptr {
		if val != "" {
			result = append(result, val)
		}
	}
	*ptr = result
	return len(result)
}
