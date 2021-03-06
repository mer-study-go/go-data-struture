package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	}
	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	fmt.Println(n.value)
}
