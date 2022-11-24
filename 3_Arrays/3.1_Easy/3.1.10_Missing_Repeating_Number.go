// Approach 1
// TC & SC - N, N
func repeat_missing(nums []int) []int {
	count := make([]int, len(nums))
	for _, v := range nums {
		count[v-1] += 1
	}
	res := make([]int, 2)
	for i, v := range count {
		if v == 0 {
			res[1] = i + 1
		} else if v == 2 {
			res[0] = i + 1
		}
	}
	return res
}

// Approach 2
// Refer Notion - Dry run example
// TC & SC - N, 1
// But squared might not be possible all the times
func repeat_missing2(nums []int) []int {
	n := len(nums)
	s := (n * (n + 1)) / 2
	s2 := (n * (n + 1) * ((2 * n) + 1)) / 6

	s_1 := 0
	s_2 := 0
	for _, v := range nums {
		s_1 += v
		s_2 += (v * v)
	}

	a1 := s - s_1
	a2 := s2 - s_2

	x_plus_y := a2 / a1
	x := (a1 + x_plus_y) / 2
	y := x - a1

	return []int{y, x}

}

// Approach 3
// Refer Notion - Dry run example
// TC & SC - N, 1
// Tricky
func repeat_missing3(nums []int) []int {
	var xor1 int
	for _, v := range nums {
		xor1 ^= v
	}
	for i := 1; i <= len(nums); i++ {
		xor1 ^= i
	}

	set_bit := xor1 & ^(xor1 - 1)

	var bucket1, bucket2 int
	for _, v := range nums {
		if (set_bit&v)&1 == 1 {
			bucket1 ^= v
		} else {
			bucket2 ^= v
		}
	}
	for i := 1; i <= len(nums); i++ {
		if (set_bit&i)&1 == 1 {
			bucket1 ^= i
		} else {
			bucket2 ^= i
		}
	}

	// Since numbers can be swapped, do a check
	isRepeatingNumber := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == bucket1 {
			isRepeatingNumber++
		}
	}

	if isRepeatingNumber == 2 {
		return []int{bucket1, bucket2}
	}
	return []int{bucket2, bucket1}

}

