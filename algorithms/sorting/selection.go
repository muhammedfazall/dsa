package main

func SelectionSort(a []int) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}

			if min != i {
				a[i], a[min] = a[min], a[i]
			}
		}

	}
}

// func main() {

// 	a := []int{4, 7, 1, 0, 8, 2, 6, 3}

// 	SelectionSort(a)

// 	fmt.Println(a)
// }
