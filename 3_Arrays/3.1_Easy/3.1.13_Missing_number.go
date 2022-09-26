// Approach 1
func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

// Approach 2
func missingNumber(nums []int) int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}
	return xor
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 4M
// Time Implement - 10M