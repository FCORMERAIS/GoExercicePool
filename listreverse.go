package piscine

func ListReverse(l *List) {
	unMaillon := &List{Head: nil}
	for i := 0; i < ListSize(l); i++ {
		ListPushFront(unMaillon, removeLast(l))
	}
	l = unMaillon
}

func removeLast(l *List) interface{} {
	for l.Head != nil {
		l.Head = l.Head.Next
	}
	result := l.Head
	l.Tail = nil
	return result
}
