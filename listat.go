package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Data == nil {
		return nil
	}
	for i := 0; i < pos; i++ {
		if l.Data == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
