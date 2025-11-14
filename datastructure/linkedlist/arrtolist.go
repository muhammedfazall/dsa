package main

import "fmt"

type Node struct {
	data int
	next *Node
	
}

func ArrayToLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{data: arr[0]}

	cur := head

	for i := 1; i < len(arr); i++ {
		newNode := &Node{data: arr[i]}
		cur.next = newNode
		cur = newNode
	}
	return head
}

func PrintList(head *Node) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Printf("%d -> ", cur.data)
	}
	fmt.Println("nil")
}

func main() {
	arr := []int{10, 20, 30}

	fmt.Println(arr)
	head := ArrayToLinkedList(arr)
	PrintList(head)
}
