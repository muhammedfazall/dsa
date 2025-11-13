package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (list *LinkedList) AddToFront(data int) {
// 	newNode := &Node{data: data}
// 	newNode.next = list.head
// 	list.head = newNode
// }

// func (list *LinkedList) AddToEnd(data int) {
// 	newNode := &Node{data: data}

// 	if list.head == nil {
// 		list.head = newNode
// 		return
// 	}

// 	current := list.head
// 	for current.next != nil {
// 		current = current.next
// 	}

// 	current.next = newNode
// }

// func (list *LinkedList) Display() {
// 	if list.head == nil {
// 		fmt.Println("list is empty")
// 	}

// 	current := list.head

// 	for current != nil {
// 		fmt.Printf("%d -> ", current.data)
// 		current = current.next
// 	}

// 	fmt.Println("nil")
// }

// func (list *LinkedList) DeleteByValue(data int) {
// 	if list.head == nil {
// 		fmt.Println("theres nothing to delete!")
// 		return
// 	}

// 	if list.head.data == data {
// 		list.head = list.head.next
// 		fmt.Println("Deleted", data)
// 		return
// 	}

// 	current := list.head
// 	for current.next != nil && current.next.data != data {
// 		current = current.next
// 	}
// 	if current.next == nil {
// 		fmt.Println("Value not found in the list!:", data)
// 		return
// 	}
// 	current.next = current.next.next
// 	fmt.Println("deleted:", data)

// }

// // func main() {
// // 	list := LinkedList{}

// // 	list.AddToEnd(10)
// // 	list.AddToEnd(50)
// // 	list.AddToEnd(60)

// // 	list.AddToFront(05)
// // 	list.AddToFront(90)
// // 	list.AddToFront(40)
// // 	list.Display()

// // 	list.DeleteByValue(05)
// // 	list.Display()

// // 	list.DeleteByValue(60)
// // 	list.Display()
// // }
