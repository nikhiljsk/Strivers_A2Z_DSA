func setZeroes(nums [][]int) {
	R, C := len(nums), len(nums[0])
	var col_zero bool
	for i := 0; i < R; i++ {
		if nums[i][0] == 0 {
			col_zero = true
		}
		for j := 1; j < C; j++ {
			if nums[i][j] == 0 {
				nums[i][0] = 0
				nums[0][j] = 0
			}
		}
	}

	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if (nums[i][0] == 0) || (nums[0][j] == 0) {
				nums[i][j] = 0
			}
		}
	}

	if nums[0][0] == 0 {
		for i := 0; i < C; i++ {
			nums[0][i] = 0
		}
	}

	if col_zero {
		for i := 0; i < R; i++ {
			nums[i][0] = 0
		}
	}

}