package main

import (
	"fmt"
)

func reverseArr(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	arr[left], arr[right] = arr[right], arr[left]
	return reverseArr(arr, left+1, right-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Reverse:", reverseArr(arr, 0, len(arr)-1))
}

// Time Complexity
// Time - O(n)
// Storage - O(1)
