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
	nL := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nL
		l.Tail = nL
	} else {
		l.Tail.Next = nL
		l.Tail = nL
	}
}
