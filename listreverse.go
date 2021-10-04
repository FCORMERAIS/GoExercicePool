package piscine

func ListReverse(l *List) {
	unMaillon := &List{Head: nil, Tail := nil}
	for i := 0; i < ListSize(l); i++ {
		ListPushFront(unMaillon, removeLast(l))
	}
	l = unMaillon
}

func removeLast(l *List) interface{} {
	for l.Head != nil {
		l.Head = l.Tail
	}
	result := l.Head
	l.Head = nil
	return result
}
