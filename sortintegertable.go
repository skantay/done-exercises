package piscine

func SortIntegerTable(table []int) {
	lenght := len(table)

	if lenght <= 1 {
		return
	}

	// Bubble sort
	for i := 0; i < lenght-1; i++ {
		for j := 0; j < lenght-i-1; j++ {
			// Esli table[j] bloshe table[J+1](to est sledushi) togda...
			if table[j] > table[j+1] {
				// ...Menyaem mestami
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
