// Approach 1
func findPivot(nums []int) int {
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
	return l
}

func bSearch(nums []int, target int) (bool, int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true, mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false, l
}

func search(nums []int, target int) int {
	pivot := findPivot(nums)
	fmt.Println("Pivot", pivot)
	isFound, indx := bSearch(nums[:pivot], target)
	if isFound {
		return indx
	}
	isFound, indx = bSearch(nums[pivot:], target)
	if isFound {
		return indx + pivot
	}
	return -1
}

// Approach 2
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
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
	return -1
}