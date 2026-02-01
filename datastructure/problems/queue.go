package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type QueueList struct {
	head *Node
	tail *Node
}

func NewQueueList() *QueueList {
	return &QueueList{
		head: nil,
		tail: nil,
	}
}

func (ql *QueueList) Enqueue(data int) {

	newNode := &Node{data: data, next: nil}

	if ql.tail == nil {
		ql.head = newNode
		ql.tail = newNode
		return
	}

	ql.tail.next = newNode
	ql.tail = newNode
}

func (ql *QueueList) Dequeue() (int, error) {
	if ql.head == nil {
		return 0, errors.New("Empty!")
	}

	v := ql.head.data
	ql.head = ql.head.next
	if ql.head == nil{
		ql.tail = nil
	}
	return v,nil
}

// func (q *Queue) Enqueue(v int){
// 	q.data = append(q.data, v)
// }

// func (q *Queue) Dequeue() (int,bool) {
// 	if q.data == nil{
// 		return 0,false
// 	}
// 	v := q.data[0]
// 	q.data = q.data[1:]
// 	return v,true
// }

func main() {
	ql := NewQueueList()

	ql.Enqueue(10)
	ql.Enqueue(40)
	ql.Enqueue(50)
	ql.Enqueue(70)

	fmt.Println(ql.head.data)
	ql.Dequeue()
	fmt.Println(ql.head.data)
	ql.Dequeue()
	fmt.Println(ql.head.data)
	ql.Dequeue()
	fmt.Println(ql.head.data)

}