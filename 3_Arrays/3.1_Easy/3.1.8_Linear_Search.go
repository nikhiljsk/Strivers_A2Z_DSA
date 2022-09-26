package main

import (
	"fmt"
)

func search(nums []int, n int) int {
	for i, v := range nums {
		if v == n {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	n := 4
	fmt.Println("search:", search(arr, n))
}
