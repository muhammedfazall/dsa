package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) AddToFront(data int) {
	newNode := &Node{data: data}

	newNode.next = list.head
	newNode.prev = nil

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	list.head.prev = newNode
	list.head = newNode
}

func (list *LinkedList) AddToEnd(data int) {
	newNode := &Node{data: data}

	newNode.prev = list.tail
	newNode.next = nil

	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	list.tail.next = newNode
	list.tail = newNode
}

func (list *LinkedList) Display() {
	for current := list.head; current != nil; current = current.next {
		fmt.Printf("%d",current.data)
		if current.next != nil {
			fmt.Printf("<->")
		}
	}
	fmt.Println()
}

func (l * LinkedList) DisplaReverse() {
	for cur := l.tail; cur != nil ; cur = cur.prev{
		fmt.Printf("%d",cur.data)
		if cur.prev != nil {
			fmt.Printf("<->")
		}
	}
	fmt.Println()
}

func main()  {
	l := LinkedList{}

	l.Display()

	l.AddToFront(50)

	l.Display()

	l.AddToEnd(40)

	l.DisplaReverse()
	l.Display()
}