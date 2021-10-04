package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos == 0 {
		return l
	}
	for i := 0; i <= pos+1; i++ {
		l.Data = l.Next
		if l.Next == nil && i <= pos {
			return nil
		}
	}
	return l.Next
}
