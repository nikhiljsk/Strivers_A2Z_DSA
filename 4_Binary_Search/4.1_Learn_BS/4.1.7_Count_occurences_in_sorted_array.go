package main

import "fmt"

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

func searchRangeCount(nums []int, target int) int {
	count := searchRange(nums, target)
	if count[0] == -1 {
		return 0
	}
	return count[1] - count[0] + 1
}

func main() {
	arr := []int{1, 2, 2, 2, 2, 3}
	target := 2
	fmt.Println(searchRangeCount(arr, target))
}
