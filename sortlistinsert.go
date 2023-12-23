package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	new := &NodeI{Data: data_ref, Next: l}

	ListSort(new)

	return new
}
