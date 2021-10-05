package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var currentNode *NodeI = n1
	if n1 != nil && n1 != nil {
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = n2
	} else if n2 != nil && n1 == nil {
		currentNode = n2
	} else if n1 != nil {
	}
	currentNode = n1
	if currentNode == nil {
		return n1
	} else {
		for i := 0; i < 52; i++ {
			for i := 0; currentNode.Next != nil; i++ {
				if currentNode.Data > currentNode.Next.Data {
					currentNode.Data, currentNode.Next.Data = currentNode.Next.Data, currentNode.Data
				}
				currentNode = currentNode.Next
			}
			currentNode = n1
		}
	}
	return n1
}
