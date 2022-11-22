func findFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	prevMid := -1
	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] == target {
			prevMid = mid
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return prevMid
}

func findLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	prevMid := -1
	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] == target {
			prevMid = mid
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return prevMid
}

func searchRange(nums []int, target int) []int {
	mi := findFirst(nums, target)
	ma := findLast(nums, target)
	return []int{mi, ma}
}