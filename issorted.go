package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var ascending bool = true
	var descending bool = true

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				ascending = false
			} else if f(a[i], a[j]) < 0 {
				descending = false
			}
		}
	}

	if ascending == true {
		return true
	}

	return descending
}
