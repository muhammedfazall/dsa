package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Delete(root *Node, value int) *Node{
	if root == nil{
		return nil
	}

	if value < root.Value{
		root.Left = Delete(root.Left,value)
	}else if value > root.Value{
		root.Right = Delete(root.Right,value)
	}else{

		if root.Left == nil && root.Right == nil{
		 	return nil
		}

		if root.Left == nil{
			return root.Right
		}

		if root.Right == nil{
			return root.Left
		}

		successor := findMin(root.Right)
		root.Value = successor.Value
		root.Right = Delete(root.Right,successor.Value)

	}
	return root
}

func findMin(root *Node) *Node {
	for root.Left != nil{
		root = root.Left
	}
	return root
}

func Insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{Value: value}
	}

	if root.Value > value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}
	return root
}

func Search(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if value == root.Value {
		return true
	} else if value < root.Value {
		return Search(root.Left, value)
	} else {
		return Search(root.Right, value)
	}
}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.Value, " ")
	Inorder(root.Right)
}

func Preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.Value, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}

	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.Value, " ")
}

func main() {
	var tree *Node

	tree = Insert(tree, 50)
	tree = Insert(tree, 20)
	tree = Insert(tree, 70)

	Inorder(tree)
}
