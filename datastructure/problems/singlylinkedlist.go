package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (l *LinkedList) addfront(data int) {
// 	newNode := &Node{
// 		data: data,
// 		next: l.head,
// 	}
// 	l.head = newNode
// }

// func (l *LinkedList) addback(data int) {
// 	newNode := &Node{data: data}

// 	if l.head == nil {
// 		l.head = newNode
// 		return
// 	}

// 	current := l.head

// 	for current.next != nil {
// 		current = current.next
// 	}
// 	current.next = newNode
// }

// func (l *LinkedList) display() {
// 	for current := l.head; current != nil; current = current.next {
// 		fmt.Printf("%d -> ",current.data)
// 	}
// 	fmt.Println("nil")
// }

// func main() {
// 	list := LinkedList{}
// 	list.display()

// 	list.addback(10)
// 	list.addfront(20)
// 	list.display()

// }
