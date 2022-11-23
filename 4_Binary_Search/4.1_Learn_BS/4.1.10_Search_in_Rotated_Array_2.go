func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}

		// The only difference from sorted array search I
		if (nums[mid] == nums[l]) && (nums[mid] == nums[r]) {
			l++
			continue
		}

		if nums[l] <= nums[mid] {
			if (nums[l] <= target) && (nums[mid] >= target) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if (nums[mid+1] <= target) && (nums[r] >= target) {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}