package piscine

func StrRev(s string) string {
	var reversed string

	for _, v := range s {
		reversed = string(v) + reversed
	}

	return reversed
}
