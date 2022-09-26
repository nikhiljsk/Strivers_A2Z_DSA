package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	freq := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if _, ok := freq[arr[i]]; ok {
			freq[arr[i]] += 1
		} else {
			freq[arr[i]] = 1
		}
	}
	fmt.Println("Least & Max:")
	minNum, maxNum := 0, 0
	min, max := 2147483647, -2147483648
	for k, v := range freq {
		if v < min {
			minNum = k
			min = v
		}
		if v > max {
			maxNum = k
			max = v
		}
	}
	fmt.Println("Min -> ", minNum, ": ", min)
	fmt.Println("Max -> ", maxNum, ": ", max)
}
