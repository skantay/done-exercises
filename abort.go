package piscine

func Abort(a, b, c, d, e int) int {
	sort := []int{a, b, c, d, e}
	for i := 0; i < len(sort)-1; i++ {
		for j := 0; j < len(sort)-i-1; j++ {
			if sort[j] > sort[j+1] {
				sort[j], sort[j+1] = sort[j+1], sort[j]
			}
		}
	}
	return sort[2]
}
