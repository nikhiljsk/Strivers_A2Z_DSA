func singleNonDuplicate(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

// The property is that, all elements (until you find the not repeating element) will take even and right next to it places
// Once the single element is found, it will start taking odd place and next to it.
func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if (mid%2 == 0 && nums[mid] == nums[mid+1]) || (mid%2 == 1 && nums[mid] == nums[mid-1]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// https://leetcode.com/problems/single-element-in-a-sorted-array/solutions/627921/java-c-python3-easy-explanation-o-logn-o-1/?orderBy=most_votes