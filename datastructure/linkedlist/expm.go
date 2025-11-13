package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (list *LinkedList) AddFro(data int) {
// 	newNode := &Node{data: data}
// 	newNode.next = list.head
// 	list.head = newNode
// }


// func (list *LinkedList) Display() {
	
// 	for current := list.head; current != nil; current = current.next {
// 		fmt.Printf("%d", current.data)
// 		if current.next != nil {
// 			fmt.Print(" -> ")
// 		}
// 	}
// 	fmt.Println()
// }

// // func (list *LinkedList) Display() {
// // if list.head == nil {
// // 		fmt.Println("list is empty!")
// // 		return
// // 	}
// // 	current := list.head
// // 	fmt.Println("linkedlist : ")
// // 	for current != nil {
// // 		fmt.Printf("%d -> ",current.data)
// // 		current = current.next
// // 	}
// // 	fmt.Println("nil")
// // }

// func (list *LinkedList) AddEnd(data int) {
// 	newNode := &Node{data:data}

// 	if list.head == nil{
// 		list.head = newNode
// 		return
// 	}

// 	current := list.head
	
// 	for current.next != nil {
// 		current = current.next
// 	}
// 	current.next = newNode
// }

// func (list *LinkedList) DeleteByV(data int) {
// 	if list.head == nil {
// 		fmt.Println("Nothing to delete in the list")
// 		return
// 	}

// 	if list.head.data == data{
// 		list.head = list.head.next
// 		fmt.Println("Deleted",data)
// 		return
// 	}

// 	current := list.head

// 	for current.next != nil && current.next.data != data{
// 		current = current.next
// 	}
// 	if current.next == nil{
// 		fmt.Println("Not found the value",data)
// 		return
// 	}
// 	current.next = current.next.next
// 	fmt.Println("Deleted",data)

// }

// func main() {
// 	list := LinkedList{}

// 	// list.AddFro(80)
// 	// list.AddFro(70)
// 	// list.AddFro(50)

// 	// list.AddEnd(100)
// 	// list.AddEnd(20)
// 	// list.AddEnd(10)

// 	list.Display()

// }
