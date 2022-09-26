// Approach 1
func reverse(nums []int, start, end int) []int {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// Approach 2
func rotate(nums []int, k int) {
	var i, curr, count, n, next_pos, j int
	n = len(nums)
	k %= n
	for count < n {
		next_pos = (i + k) % n
		curr, nums[next_pos] = nums[next_pos], nums[i]
		count++
		j = next_pos
		for count < n && j != i {
			next_pos = (j + k) % n
			curr, nums[next_pos] = nums[next_pos], curr
			count++
			j = next_pos
		}
		i++
	}
}

// Other approaches
// Can be copied to a new array directly
// Using GCD? Divide into sets, and rotate once in each set

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 40M
// Time Implement - 15M