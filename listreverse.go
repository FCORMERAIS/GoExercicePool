package piscine

func ListReverse(l *List) {
	unMaillon := &List{Head: nil, Tail: nil}
	for i := 0; i < ListSize(l); i++ {
		ListPushFront(unMaillon, removeLast(l))
	}
	l = unMaillon
}

func removeLast(l *List) interface{} {
	tmp := l
	for tmp.Head.Next != nil {
		tmp.Head = tmp.Head.Next
	}
	result := tmp.Head
	tmp.Head = nil
	return result
}
