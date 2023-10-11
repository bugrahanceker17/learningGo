package main

import "fmt"

type Node struct {
	next *Node
	data MyItem
}

type LinkedList struct {
	head *Node
}

type MyItem int

func main() {
	linkedList := NewLinkedList()
	linkedList.insert(1)
	linkedList.insert(2)
	linkedList.insert(3)
	linkedList.insert(4)
	linkedList.insert(5)
	linkedList.insert(6)

	linkedList.delete(3)

	linkedList.print()
}

func (l *LinkedList) print() {
	curr := l.head

	for curr != nil {
		fmt.Printf(" -> %d", curr.data)
		curr = curr.next
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) insert(data MyItem) {
	n := &Node{
		next: nil,
		data: data,
	}

	if l.head == nil {
		l.head = n
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = n
}

func (l *LinkedList) delete(value MyItem) {
	curr := l.head
	prev := l.head

	for curr != nil {
		if curr.data == value {
			prev.next = curr.next
		}

		prev = curr
		curr = curr.next
	}
}
