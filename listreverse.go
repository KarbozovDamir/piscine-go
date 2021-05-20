package piscine

func ListReverse(l *List) {
	y := l.Head
	l.Tail = l.Head
	var x, z *NodeL
	for y != nil {
		x = y.Next
		y.Next = x
		z = y
		y = z
	}
	l.Head = x
}
