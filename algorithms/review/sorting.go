package main

import "fmt"

// import "fmt"

// func BubbleSort(a []int) {
// 	n := len(a)
// 	swapped := true

// 	for swapped {
// 		swapped = false

// 		for i := 0; i < n-1; i++ {
// 			if a[i] > a[i+1] {
// 				a[i], a[i+1] = a[i+1], a[i]
// 				swapped = true
// 			}
// 		}
//      n--
// 	}
// }

// func InsertionSort(a []int) {
// 	n := len(a)

// 	for i := 1; i < n; i++ {
// 		j := i - 1
// 		temp := a[i]

// 		for j >= 0 && a[j] > temp {
// 			a[j+1] = a[j]
// 			j--
// 		}

// 		a[j+1] = temp
// 	}
// }

// func SelectionSort(a []int) {
// 	n := len(a)
// 	for i := 0; i < n-1; i++ {
// 		min := i
// 		for j := i + 1; j < n; j++ {
// 			if a[j] < a[min] {
// 				min = j
// 			}
// 			if min != i {
// 				a[j], a[min] = a[min], a[i]
// 			}
// 		}
// 	}
// }

// func main() {
// 	a := []int{9, 7, 5, 4, 3, 1}
// 	b := []int{9, 7, 5, 4, 3, 1}
// 	c := []int{9, 7, 5, 4, 3, 1}

// 	BubbleSort(a)
// 	InsertionSort(b)
// 	SelectionSort(c)

// 	fmt.Println("bubble:" ,a)
// 	fmt.Println("Insertion:",b)
// 	fmt.Println("selection:",c)
// }

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{Data: data}
	}

	if data < root.Data {
		root.Left = Insert(root.Left, data)
	} else {
		root.Right = Insert(root.Right, data)
	}

	return root
}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.Data, " ")
	Inorder(root.Right)
}

func Preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.Data, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}

	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.Data, " ")
}

func main() {

	var bst *Node

	bst = Insert(bst, 50)
	bst = Insert(bst, 100)
	bst = Insert(bst, 40)

	Inorder(bst)
	fmt.Println()
	Preorder(bst)
	fmt.Println()
	Postorder(bst)
}
