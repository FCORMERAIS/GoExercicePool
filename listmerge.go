package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil {
		l1 = l2
	} else if l2 == nil{
		l2 = l1
	} else {
		l1.Tail.Next = l2.Head
		tmp := l1
		for l1.Head.Next != nil {
			tmp.Head = tmp.Head.Next
		}
		l1.Tail = tmp.Head
	}
}
