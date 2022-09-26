package main

import (
	"fmt"
)

func maxArr(arr []int) int {
	max := -2147483648
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println(maxArr(arr))
}
