package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listPushBack2(l, data_ref)
	ListSort(l)
	return l
}

func listPushBack2(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
