func removeDuplicates(nums []int) int {
	prev := 0
	for curr := 1; curr < len(nums); curr++ {
		if nums[prev] != nums[curr] {
			nums[prev+1] = nums[curr]
			prev++
		}
	}
	return prev + 1 // since index
}