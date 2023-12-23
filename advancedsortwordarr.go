package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			num := f(string(a[j]), string(a[j+1]))
			if num == 1 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
