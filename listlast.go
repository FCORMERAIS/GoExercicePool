package piscine

func ListLast(l *List) interface{} {
	for l.Head.Next != nil {
		l.Head = l.Head.Next
	}
	return &l.Head
}
