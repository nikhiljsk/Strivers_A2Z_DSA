func nextPermutation(nums []int) {
	i, j := len(nums)-1, len(nums)-1

	// Find the descending element
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	// If no descending element, just reverse
	if i == 0 {
		reverse(nums, 0)
		return
	}

	// Find the just highest number in the sub-array
	k := i - 1
	for nums[k] >= nums[j] {
		j--
	}
	// Swap and reverse
	nums[k], nums[j] = nums[j], nums[k]
	reverse(nums, k+1)
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}