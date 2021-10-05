package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil {
		l1 = l2
	} else {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail.Next
	}
}
