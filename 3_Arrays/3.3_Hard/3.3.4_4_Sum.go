func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			l, r := j+1, n-1
			target2 := target - nums[i] - nums[j]
			for l < r {
				curr_sum := nums[l] + nums[r]
				if curr_sum < target2 {
					l++
				} else if curr_sum > target2 {
					r--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					// Skip duplicates for 3rd
					for (l < r) && (nums[l+1] == nums[l]) {
						l++
					}

					// Skip duplicates for 4th
					for (l < r) && (nums[r-1] == nums[r]) {
						r--
					}
					l++
					r--
				}
			}
			// Skip duplicates for 2nd
			for (j+1 < n) && (nums[j+1] == nums[j]) {
				j++
			}
		}
		// Skip duplicates for 1st
		for (i+1 < n) && (nums[i+1] == nums[i]) {
			i++
		}
	}
	return res
}