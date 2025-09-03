package solution

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	// 先頭の要素を指すポインタ
	Head *Node[T]
	// 末尾の要素を指すポインタ
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		// 戦闘要素が存在しない場合、そのまま挿入する
		l.Head = n
		l.Tail = n
		return
	}

	if pos <= 0 {
		// 負の値を指定された場合、先頭に挿入する
		n.Next = l.Head
		l.Head = n
		return
	}

	curNode := l.Head
	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			// 終端に到達した場合、末尾に挿入する
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
}

func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}

	return -1
}
