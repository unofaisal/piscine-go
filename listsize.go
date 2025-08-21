package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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

func ListSize(l *List) int {
  counter := 0
  for l.Head != nil {
    l.Head = l.Head.Next
    counter++
  }
  return counter
}