package main

import (
	"chapter08/solution"
	"errors"
	"fmt"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Double[T Numeric](val T) T {
	return val * 2
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return fmt.Sprintf("PrintableInt: %d\n", pi)
}

type PrintableFloat64 float64

func (pf PrintableFloat64) String() string {
	return fmt.Sprintf("PrintableFloat64: %f\n", pf)
}

func Print[T Printable](val T) {
	fmt.Println(val.String())
}

type Node[T comparable] struct {
	Next  *Node[T]
	Value T
}

type List[T comparable] struct {
	Nodes *Node[T]
}

func (l *List[T]) Add(val T) {
	if l.Nodes == nil {
		l.Nodes = &Node[T]{Value: val}
		return
	}

	l.Nodes.Add(val)
}

func (n *Node[T]) Add(val T) {
	if n == nil {
		n = &Node[T]{Value: val}
		return
	}

	if n.Next == nil {
		n.Next = &Node[T]{Value: val}
		return
	}

	n.Next.Add(val)
}

func (l *List[T]) Insert(val T, idx int) error {
	if idx < 0 {
		return errors.New("不正なインデックスです。")
	}

	return l.Nodes.Insert(val, idx)
}

func (n *Node[T]) Insert(val T, idx int) error {
	if idx < 0 {
		return errors.New("不正なインデックスです。")
	}

	if idx == 0 {
		cv := n.Value
		cn := n.Next
		n.Value = val
		n.Next = &Node[T]{Value: cv, Next: cn}
		return nil
	}

	if n.Next == nil {
		return errors.New("不正なインデックスです。")
	}

	return n.Next.Insert(val, idx-1)
}

func (l *List[T]) Index(val T) int {
	if l.Nodes == nil {
		return -1
	}

	return l.Nodes.Index(val, 0)
}

func (n *Node[T]) Index(val T, idx int) int {
	if n == nil {
		return -1
	}

	if n.Value == val {
		return idx
	}

	return n.Next.Index(val, idx+1)
}

func main() {
	fmt.Println(Double(42))
	fmt.Println(Double(3.14))

	Print(PrintableInt(42))
	Print(PrintableFloat64(3.14))

	l := List[int]{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	fmt.Printf("1 is in List[%d]\n", l.Index(1))
	fmt.Printf("2 is in List[%d]\n", l.Index(2))
	fmt.Printf("3 is in List[%d]\n", l.Index(3))

	l.Insert(4, 0)
	l.Insert(5, 1)
	l.Insert(6, 2)
	fmt.Printf("1 is in List[%d]\n", l.Index(1))
	fmt.Printf("2 is in List[%d]\n", l.Index(2))
	fmt.Printf("3 is in List[%d]\n", l.Index(3))
	fmt.Printf("4 is in List[%d]\n", l.Index(4))
	fmt.Printf("5 is in List[%d]\n", l.Index(5))
	fmt.Printf("6 is in List[%d]\n", l.Index(6))

	ls := &solution.List[int]{}
	ls.Add(5)
	ls.Add(10)
	fmt.Println(ls.Index(5))
	fmt.Printf("5 is in List[%d]\n", ls.Index(5))
	fmt.Printf("10 is in List[%d]\n", ls.Index(10))
	fmt.Printf("20 is in List[%d]\n", ls.Index(20))

	ls.Insert(100, 0)
	fmt.Printf("5 is in List[%d]\n", ls.Index(5))
	fmt.Printf("10 is in List[%d]\n", ls.Index(10))
	fmt.Printf("20 is in List[%d]\n", ls.Index(20))
	fmt.Printf("100 is in List[%d]\n", ls.Index(100))

	ls.Insert(200, 1)
	fmt.Printf("5 is in List[%d]\n", ls.Index(5))
	fmt.Printf("10 is in List[%d]\n", ls.Index(10))
	fmt.Printf("20 is in List[%d]\n", ls.Index(20))
	fmt.Printf("100 is in List[%d]\n", ls.Index(200))

	fmt.Println("=========================")
	for curNode := ls.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	fmt.Println("=========================")
	ls.Insert(300, 10)
	for curNode := ls.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	fmt.Println("=========================")
	ls.Add(400)
	for curNode := ls.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	fmt.Println("=========================")
	ls.Insert(500, 6)
	for curNode := ls.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
