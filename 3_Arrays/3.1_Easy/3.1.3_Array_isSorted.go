package main

import (
	"fmt"
)

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] >= arr[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	// arr := []int{1, 4, 6, 2, 3, 5}
	arr := []int{}
	fmt.Println("isSorted:", isSorted(arr))
}
