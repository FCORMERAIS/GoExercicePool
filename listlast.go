package piscine

func ListLast(l *List) interface{} {
	tmp := l
	if tmp.Head == nil {
		return nil
	}
	for tmp.Head.Next != nil {
		tmp.Head = tmp.Head.Next
	}
	return tmp.Head.Data
}
