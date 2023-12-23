package piscine

func StrLen(s string) int {
	var len int = 0

	for _, v := range s {
		v++
		len++
	}

	return len
}
