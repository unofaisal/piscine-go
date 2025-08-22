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

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	cur := l.Head
	for cur != nil {
		if CompStr(cur.Data, data_ref) {
			if prev == nil {
				l.Head = cur.Next
			} else {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}

		cur = cur.Next
	}
}