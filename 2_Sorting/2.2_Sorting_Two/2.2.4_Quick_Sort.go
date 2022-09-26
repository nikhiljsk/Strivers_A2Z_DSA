package main

import (
	"fmt"
)

func partition(arr []int, l, r int) int {
	mid := (l + r) / 2
	pivot := arr[mid]
	i, j := l, r
	for i < j {
		for arr[i] <= pivot && i <= r-1 {
			i++
		}
		for arr[j] > pivot && j >= l-1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[mid] = arr[mid], arr[j]
	return j
}

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(arr, l, r)
	QuickSort(arr, l, i-1)
	QuickSort(arr, i+1, r)
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("After Sorting:", arr)
}
