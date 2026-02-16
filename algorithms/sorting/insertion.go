package main

func InsertionSort(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		temp := a[i]
		j := i - 1

		for j >= 0 && a[j] > temp {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = temp
	}
}
