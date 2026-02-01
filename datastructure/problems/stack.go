package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Node struct {
// 	data int
// 	next *Node
// }

// type StackList struct {
// 	head *Node
// }

// func NewStackList() *StackList {
// 	return &StackList{head: nil}
// }

// func (s *StackList) push(v int) {
// 	newNode := &Node{data: v, next: s.head}
// 	s.head = newNode
// }

// func (s *StackList) pop() (int, error) {
// 	if s.head == nil {
// 		return 0, errors.New("empty")
// 	}
// 	v := s.head.data
// 	s.head = s.head.next
// 	return v,nil
// }

// func (s *StackList) size() int {
// 	count := 0
// 	for cur := s.head; cur != nil ; cur = cur.next{
// 		count++
// 	}
// 	return count
// }

// func main() {
// 	sl := StackList{}
// 	sl.push(10)
// 	sl.push(50)

// 	fmt.Println(sl.head.data)
// }
