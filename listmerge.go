package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil {
		l1.Head = l2.Head
	} else if l2 == nil {
		l2.Head = l1.Head
	} else {
		l1.Tail.Next = l2.Head
	}
}
