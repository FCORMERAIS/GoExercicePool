package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	tmp := l
	for i := 0; i < 45; i++ {
		for tmp.Next != nil {
			if tmp.Data > tmp.Next.Data {
				tmp.Data, tmp.Next.Data = tmp.Next.Data, tmp.Data
			}
			tmp.Data = tmp.Next.Data
		}
		tmp = l
	}
	return l
}
