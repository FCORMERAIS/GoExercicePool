package piscine

func ListReverse(l *List) {
	currentNode := l.Head
	var Next *NodeL
	var previousNode *NodeL
	l.Tail = l.Head
	for currentNode != nil {
		Next, currentNode.Next = currentNode.Next, previousNode
		previousNode, currentNode = currentNode, Next
	}
	l.Head = previousNode
}
