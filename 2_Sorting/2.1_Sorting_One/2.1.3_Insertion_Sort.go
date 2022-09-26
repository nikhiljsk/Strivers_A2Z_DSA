package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	for i := 1; i < len(arr); i++ {
		j := i - 1
		curr := arr[i]
		for arr[j] > curr && j >= 0 {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
	fmt.Println("After Sorting:", arr)
}
