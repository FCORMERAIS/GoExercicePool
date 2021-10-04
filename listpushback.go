package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	unMaillon := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = unMaillon
		l.Tail = unMaillon
	} else {
		l.Tail.Next = unMaillon
		l.Tail = unMaillon
	}
}
