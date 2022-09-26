package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println("After Sorting:", arr)
}
