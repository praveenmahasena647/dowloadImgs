package lists

type Node struct {
	Prev *Node
	Name string
	Link string
	Next *Node
}

func (l *List) Add(n Node) {
	if l.Head == nil {
		l.Head, l.Tail = &n, &n
		return
	}
	l.Tail.Next, n.Prev, l.Tail = &n, l.Tail, &n
}
