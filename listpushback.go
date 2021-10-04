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
	unMaillon := &NodeL{Data : data}
	if l.Head == nil {
		l.Head = unMaillon
		l.Tail = nil
	}else {
		unMaillon.Next = l.Head
		l.Head = unMaillon
	}
}
