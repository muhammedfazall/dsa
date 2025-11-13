package main

import (
	"fmt"
)

type DNode struct {
	data int
	next *DNode
	prev *DNode
}

type DlinkedList struct {
	head *DNode
	tail *DNode
}

func (dlist *DlinkedList) AddToFronttt(data int) {
	newNode := &DNode{data: data}

	newNode.next = dlist.head
	newNode.prev = nil

	if dlist.head == nil {
		dlist.head = newNode
		dlist.tail = newNode
		return
	}

	dlist.head.prev = newNode
	dlist.head = newNode
}

func (dlist *DlinkedList) Ddisplay() {
	for cur := dlist.head; cur != nil; cur = cur.next {
		fmt.Printf("%d", cur.data)
		if cur.next != nil {
			fmt.Print(" <-> ")
		}
	}
	fmt.Println()
}

func (dlist *DlinkedList) AddToEnddd(data int) {
	newNode := &DNode{data: data}

	newNode.prev = dlist.tail
	newNode.next = nil

	if dlist.tail == nil {
		dlist.head = newNode
		dlist.tail = newNode
		return
	}

	dlist.tail.next = newNode
	dlist.tail = newNode
}

func (dlist *DlinkedList) DisplayReverse() {
	for cur := dlist.tail; cur != nil; cur = cur.prev {
		fmt.Printf("%d", cur.data)
		if cur.prev != nil {
			fmt.Print(" <-> ")
		}
	}
	fmt.Println()

}

func (dlist *DlinkedList) DeletebyValuee(data int) {
	if dlist.head == nil {
		fmt.Println("nothing to delete!")
		return
	}

	cur := dlist.head

	for cur != nil && cur.data != data {
		cur = cur.next
	}

	if cur == nil {
		fmt.Println("Not found!")
		return
	}

	if cur.prev == nil {
		dlist.head = cur.next
	} else {
		cur.prev.next = cur.next
	}

	if cur.next == nil {
		dlist.tail = cur.prev
	} else {
		cur.next.prev = cur.prev
	}

}

func main() {
	dlist := DlinkedList{}

	dlist.AddToFronttt(40)
	dlist.AddToFronttt(20)
	dlist.Ddisplay()
	dlist.AddToEnddd(70)
	dlist.AddToEnddd(66)
	dlist.AddToEnddd(42)
	dlist.Ddisplay()
	dlist.DeletebyValuee(20)
	dlist.Ddisplay()
	dlist.DisplayReverse()

}
