package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {

			if a[i] == -777 {
				break
			}

			if a[i] == a[j] && i != j {
				a[i] = -777
				a[j] = -777
				break
			}

			if j == len(a)-1 {
				return a[i]
			}
		}
	}

	return -1
}
