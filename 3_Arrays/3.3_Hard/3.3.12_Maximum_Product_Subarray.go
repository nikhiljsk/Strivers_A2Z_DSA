func max(arr []int) int {
	maxi := -11
	for _, v := range arr {
		if maxi < v {
			maxi = v
		}
	}
	return maxi
}

func min(arr []int) int {
	mini := 11
	for _, v := range arr {
		if mini > v {
			mini = v
		}
	}
	return mini
}

func maxProduct(nums []int) int {
	res, imax, imin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}

		candidates := []int{nums[i], imax * nums[i], imin * nums[i]}
		imax = max(candidates)
		imin = min(candidates)

		res = max([]int{imax, res})
	}
	return res
}