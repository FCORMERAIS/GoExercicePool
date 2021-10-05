package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	tmp := l
	if tmp.Head != nil {
		if tmp.Head == data_ref {
			unMaillon := &NodeL{Data: nil, Next: nil}
			unMaillon = l.Head
			l.Head = l.Head.Next
			unMaillon.Next = nil
		}
	}
	if tmp.Head.Next != nil {
		for tmp != nil {
			if tmp.Head.Next == data_ref {
				unMaillon := &NodeL{Data: nil, Next: nil}
				unMaillon = l.Head
				l.Head = l.Head.Next
				unMaillon.Next = nil
			}
			tmp.Head = l.Head.Next
		}
	}
}
