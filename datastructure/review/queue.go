package main

import (
	"fmt"
)

type SliceQueue struct {
	items []int
}

func (q *SliceQueue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *SliceQueue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func main() {

	qu := SliceQueue{}

	qu.Enqueue(10)
	qu.Enqueue(50)
	qu.Enqueue(60)
	fmt.Println(qu.items)

	v, _ := qu.Dequeue()
	fmt.Println("dequeued:", v)
	fmt.Println(qu.items)
}
