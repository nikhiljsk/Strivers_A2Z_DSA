// Approach 1
// Kadane's Algorithm
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
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
	start, end := 0, 0
	for i, v := range nums {
		localMax = max(v, localMax+v)
		if globalMax < localMax {
			globalMax = localMax
			end = i
		}
		if localMax < 0 {
			start = i + 1
		}
	}
	fmt.Println(start, end)
	return globalMax
}