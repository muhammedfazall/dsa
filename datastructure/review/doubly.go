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

func (l *LinkedList) AddtoFront(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode

}

func (l *LinkedList) AddtoEnd(data int) {
	newNode := &Node{data: data}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) Display() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}


	for cur := l.head ; cur != nil ; cur = cur.next{
		fmt.Println(cur.data)
		
	}
}

// func main() {
// 	list := LinkedList{}

// 	list.AddtoFront(10)
// 	list.AddtoFront(20)
// 	list.AddtoEnd(40)
// 	list.AddtoEnd(80)
// 	list.Display()
// }
