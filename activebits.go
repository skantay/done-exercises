package piscine

func ActiveBits(n int) int {
	var count int

	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
