package piscine

func ListPushFront(l *List, data interface{}) {
	unMaillon := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = unMaillon
		l.Tail = unMaillon
	} else {
		unMaillon.Next = l.Head
		l.Head = unMaillon
	}
}
