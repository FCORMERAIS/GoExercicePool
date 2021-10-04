package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var tmp *NodeL = l.Head
	if tmp.Data == nil {
		return nil
	}
	for tmp != nil {
		if comp(tmp, ref) == true {
			return &tmp.Data
		}
		tmp = tmp.Next
	}
	return nil
}
