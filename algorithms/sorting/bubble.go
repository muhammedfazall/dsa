package main

// // 'while' style

func BubbleSort(a []int) {
	n := len(a)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		n--
	}
}

// func main() {

// 	a := []int{4, 7, 1, 0, 8, 2, 6, 3}

// 	BubbleSort(a)

// 	fmt.Println(a)
// }

// using range (idiomatic)

// func BubbleSort(a []int) {
// 	n := len(a)

// 	for i := 0; i < n-1; i++ {
// 		// We use a slice of the array to limit the range
// 		for j := range a[:n-1-i] {
// 			if a[j] > a[j+1] {
// 				a[j], a[j+1] = a[j+1], a[j]
// 			}
// 		}
// 	}
// }

// func BubbleSort(a []int) {
// 	n := len(a)
// 	for i := 0; i < n-1; i++ {
// 		swapped := false
// 		for j := 0; j < n-i-1; j++ {
// 			if a[j] > a[j+1] {
// 				a[j], a[j+1] = a[j+1], a[j]
// 				swapped = true
// 			}
// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// }
