package main

import "fmt"

// Approach 1
// Kadane's Algorithm
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray2(nums []int) int {
	localMax := -1000000
	globalMax := -1000000
	for _, v := range nums {
		localMax = max(v, localMax+v)
		globalMax = max(globalMax, localMax)
	}
	return globalMax
}

// Approach 2
// If asked to print the subarray
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	localMax := -1000000
	globalMax := -1000000
	start := 0
	var res [2]int
	for i, v := range nums {
		localMax = max(v, localMax+v)
		if globalMax < localMax {
			globalMax = localMax
			// Store the result
			res[0] = start
			res[1] = i
		}
		if localMax < 0 {
			start = i + 1
		}
	}
	fmt.Println(res[0], res[1])
	return globalMax
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
