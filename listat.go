package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	truc := &List{Head: l.Next}
	if ListSize(truc) <= pos {
		return nil
	}
	if pos == 0 {
		return l
	}
	for i := 0; i <= pos; i++ {
		l.Data = l.Next
		if l.Next == nil && i+1 != pos {
			return nil
		}
	}
	return l.Next
}
