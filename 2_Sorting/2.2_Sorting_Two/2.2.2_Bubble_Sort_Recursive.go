package main

import (
	"fmt"
)

func BubbleSort(arr []int, i, n int) {
	if i >= n {
		return
	}
	for j := 0; j < len(arr)-i-1; j++ {
		if arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	BubbleSort(arr, i+1, n)
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	BubbleSort(arr, 0, len(arr)-1)
	fmt.Println("After Sorting:", arr)
}
