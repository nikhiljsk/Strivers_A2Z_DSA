// Finding min is basically the same as find the pivot point
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	last := nums[r]
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > last {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}