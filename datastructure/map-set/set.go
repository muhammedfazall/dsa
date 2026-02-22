package main

import "fmt"

type Set map[int]bool

func (s Set) Add(v int) {
	s[v] = true
}

func (s Set) Remove(v int) {
	delete(s, v)
}

func (s Set) Contains(v int) bool {
	return s[v]
}

func main() {
	s := Set{}

	s.Add(10)
	s.Add(20)
	s.Add(30)
	s.Add(10)

	fmt.Println(s)

	s.Remove(20)
	fmt.Println(s)

	fmt.Println(s.Contains(30))
}