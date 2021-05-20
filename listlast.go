package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		return nil
	}
	return l.Head.Data
}
