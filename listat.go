package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	for l.Head != nil {
		node = l.Head
		l.Head
	}
	return node.Data
}
