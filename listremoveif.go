package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL

	if l.Head != nil {

		current := l.Head
		prev = current

		if l.Head.Data == data_ref {
			l.Head = current.Next
		}

		for current != nil {

			if current.Data == data_ref {
				if current.Next == nil {
					prev.Next = nil
					return
				}
				prev.Next = current.Next
				current = current.Next
				continue
			}

			prev = current
			current = current.Next
		}
	}
}
