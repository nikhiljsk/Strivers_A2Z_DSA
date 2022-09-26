package main

import (
	"fmt"
)

func isSortedRotated(nums []int) bool {
	var count int

	// a[i] can't be greater than a[i+1] more than once
	// Last and first element are related, hence to compare them as well
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true

}

func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println("isSortedRotated:", isSortedRotated(arr))
}
