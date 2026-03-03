package main

import "fmt"

//Merge Sort:

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2 
	
	left := MergeSort(arr[:mid]) 
	right := MergeSort(arr[mid:])

	return Merge(left,right)
}

func Merge(left []int, right []int) []int {
	result := []int{}
	i,j := 0,0

	for i < len(left) && j < len(right){
		
		if left[i] <= right[j]{
			result = append(result, left[i])
			i++
		}else{
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// func QuickSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivot := arr[len(arr)/2]

// 	less := []int{}
// 	equal := []int{}
// 	greater := []int{}

// 	for _, x := range arr {
// 		if x < pivot {
// 			less = append(less, x)
// 		} else if x == pivot {
// 			equal = append(equal, x)
// 		} else {
// 			greater = append(greater, x)
// 		}
// 	}
// 	result := append(QuickSort(less), equal...)
// 	result = append(result, QuickSort(greater)...)
	
// 	return result
// }


// Quick Sort:

func QuickSort(arr []int, low,high int) {
	if low < high{
		pindex := Partition(arr,low,high)

		QuickSort(arr, low, pindex - 1)
		QuickSort(arr, pindex + 1, high)
	}
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{7, 5, 4, 1, 10, 6, 10}
	arr2 := []int{8, 2, 4, 3, 9, 6, 10}

	sorted := MergeSort(arr)

	// sortedQuick := QuickSort(arr)
	QuickSort(arr2,0,len(arr2)-1) // recursive

	fmt.Println("merge:", sorted) 
	// fmt.Println("quick:", sortedQuick)
	fmt.Println("quick:", arr2)
}