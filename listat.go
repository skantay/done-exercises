package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var count int = 0
	for l.Next != nil {
		if count == pos {
			return l
		}

		l = l.Next
		count++
	}
	return nil
}
