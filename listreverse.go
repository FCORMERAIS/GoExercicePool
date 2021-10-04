package piscine

func ListReverse(l *List) {
	if l.Head != nil {
		unMaillon := &List{Head: nil, Tail: nil}
		for i := 0; i < ListSize(l); i++ {
			ListPushFront(unMaillon, removeLast(l))
		}
		l = unMaillon
	}
}

func removeLast(l *List) interface{} {
	tmp := l
	if tmp.Head == nil {
		return nil
	}
	if tmp.Head.Next == nil {
		return tmp.Head
	}
	for tmp.Head.Next != l.Tail {
		tmp.Head = tmp.Head.Next
	}
	result := tmp.Head.Next
	tmp.Head = tmp.Head.Next
	l.Tail = tmp.Head
	l.Tail.Next = nil
	return result
}
