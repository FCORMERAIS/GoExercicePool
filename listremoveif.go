package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var NodeL1 *NodeL = l.Head
	if l.Head != nil {
		if l.Head.Data == data_ref {
			NodeL1 = l.Head
			l.Head = l.Head.Next
			NodeL1.Next = l.Head
		}
	}
	if NodeL1 != nil {
		for NodeL1.Next != nil {
			if NodeL1.Next.Data == data_ref {
				NodeL1.Next = NodeL1.Next.Next
			}
			if NodeL1.Next != nil && NodeL1.Next.Data != data_ref {
				NodeL1 = NodeL1.Next
			}
		}
	}
}
