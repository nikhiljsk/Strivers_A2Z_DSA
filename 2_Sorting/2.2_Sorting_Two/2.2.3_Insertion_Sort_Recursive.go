package main

import (
	"fmt"
)

func InsertionSort(arr []int, i, n int) {
	if i >= n {
		return
	}
	curr := arr[i]
	j := i - 1
	for j >= 0 && arr[j] > curr {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = curr
	InsertionSort(arr, i+1, n)
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	InsertionSort(arr, 1, len(arr)-1)
	fmt.Println("After Sorting:", arr)
}
