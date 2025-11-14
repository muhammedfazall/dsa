package main

import (
	"errors"
	// "fmt"
)

type SliceStack struct {
	data []int
}

func NewStack() *SliceStack {
	return &SliceStack{data: []int{}}
}

func (s *SliceStack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *SliceStack) Pop() (int, error) {
	if s.data == nil {
		return 0, errors.New("Stack is empty")
	}
	i := len(s.data)-1
	v := s.data[i]
	s.data = s.data[:i]
	return v,nil
}

func (s *SliceStack) Peek() (int,error) {
	if s.data == nil {
		return 0, errors.New("Stack is empty")
	}
	return s.data[len(s.data)-1],nil
}

// func main() {
// 	st := NewStack()
	
// 	st.Push(50)
// 	st.Push(20)
// 	st.Push(80)
// 	fmt.Println("Stack:")
// 	fmt.Println(st.data)

// 	val,_ := st.Peek()
// 	fmt.Println("value at top:",val)

// 	v,_ := st.Pop()
// 	fmt.Println(st.data)
// 	fmt.Println("poped value:",v)
	
// }
