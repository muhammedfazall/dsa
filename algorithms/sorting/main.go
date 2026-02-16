package main

import "fmt"

func main() {

	a := []int{4, 7, 1, 0, 8, 2, 6, 3}

	// BubbleSort(a)
	// InsertionSort(a)
	SelectionSort(a)

	fmt.Println(a)
}
