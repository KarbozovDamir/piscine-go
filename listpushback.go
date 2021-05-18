package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	result := NodeL{
		Data: data,
	}
	if l.Head == nil {
		l.Head = result
	} else {
		l.Tail.Next = result
	}
	l.Tail = result
}
