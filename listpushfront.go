package piscine

func ListPushFront(l *List, data interface{}) {
	nL := &NodeL{Data: data}
	l.Tail.Next = nL
	l.Tail = nL
}
