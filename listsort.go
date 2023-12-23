package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	data := []int{}
	current := l

	for current != nil {

		data = append(data, current.Data)
		current = current.Next
	}

	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	current = l
	i := 0
	for current != nil {
		current.Data = data[i]
		i++
		current = current.Next
	}

	return l
}
