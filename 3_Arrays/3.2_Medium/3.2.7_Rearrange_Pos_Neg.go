
// Wrong Approach!!
// This will bring all the positive to last and negative to first
func rearrangeArray(nums []int) []int {
	p, n := 0, 0
	for p < len(nums)-1 && n < len(nums)-1 {
		for p < len(nums) && nums[p] < 0 {
			p++
		}
		for n < len(nums) && nums[n] > 0 {
			n++
		}
		nums[p], nums[n] = nums[n], nums[p]
	}
	return nums
}

// Approach 1
// N, N
func rearrangeArray(nums []int) []int {
	negInd, posInd := 1, 0
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res[posInd] = nums[i]
			posInd += 2
		} else {
			res[negInd] = nums[i]
			negInd += 2
		}
	}
	return res
}