package piscine

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
