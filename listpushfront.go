package piscine

func ListPushFront(l *List, data interface{}) {
	res := &NodeL{
		Data: data,
	}
	if l.Head = nil {
		l.Head = res
	} else {
		res.Next = l.Head
	}
	l.Head = res
}