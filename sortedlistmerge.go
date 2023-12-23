package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	def := n1
	for n1 != nil {
		if n1.Next == nil {
			n1.Next = n2
			break
		}
		n1 = n1.Next
	}

	ListSort(def)

	return def
}
