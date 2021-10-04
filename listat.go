package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos+1; i++ {
		l.Data = l.Next
		if l.Data == nil {
			return nil
		}
	}
	return l.Next
}
