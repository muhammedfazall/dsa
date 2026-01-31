package main

// import "errors"

// type Node struct {
// 	data int
// 	next *Node
// }

// type QueueList struct {
// 	head *Node
// 	tail *Node
// }

// func NewQueueList() *QueueList {
// 	return &QueueList{head: nil, tail: nil}
// }

// func (q *QueueList) Enqueue(v int) {
// 	newNode := &Node{data: v, next: nil}
// 	if q.tail == nil {
// 		q.head = newNode
// 		q.tail = newNode
// 		return
// 	}

// 	q.tail.next = newNode
// 	q.tail = newNode
// }

// func (q *QueueList) Dequeue() (int, error) {
// 	if q.head == nil {
// 		return 0, errors.New("empry")
// 	}
// 	val := q.head.data
// 	q.head = q.head.next

// 	if q.head == nil{
// 		q.tail = nil
// 	}
// 	return val,nil
// }

// func main() {

// }
