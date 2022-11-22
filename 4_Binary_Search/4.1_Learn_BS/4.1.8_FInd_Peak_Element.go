func findPeakElement(nums []int) int {
	// Making use of the fact that, if a random point is selected in the graphical representation of array:
	// If the point lies in a decending slope line, the peak lies to its left
	// If the point lies in a ascending slope line, the peak lies to its right
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// Tip - It is essential to use l<r here, and also r = mid. Else if you would use l<=r and r = mid-1, it would result in index out of bounds error
// Whereever you need mid+1 to be accessed, use this method

// Linear Search
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}