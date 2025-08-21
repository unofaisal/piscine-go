package piscine

func ListPushFront(l *List, data interface{}) {
	nL := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nL
		l.Tail = nL
	} else {
		nL.Next = l.Head
		l.Head = nL
	}
}
