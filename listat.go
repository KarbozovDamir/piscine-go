package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	for l.Head != nil {
		node = l.Head
		l.Head = l.Head.Next

	}
	return node.Data
}
