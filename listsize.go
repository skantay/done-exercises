package piscine

func ListSize(l *List) int {
	var count int

	if l.Head == nil {
		return 0
	}
	start := l.Head
	count++

	for start.Next != nil {
		count++
		start = start.Next
	}

	return count
}
