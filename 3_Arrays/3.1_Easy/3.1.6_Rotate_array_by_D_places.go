package main

import (
	"fmt"
)

func reverse(nums []int, start, end int) []int {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}

// Approach 1
func rotate(nums []int, k int, d string) {
	n := len(nums)
	k %= n
	if d == "right" {
		reverse(nums, 0, n-1)
		reverse(nums, 0, k-1)
		reverse(nums, k, n-1)
	} else if d == "left" {
		reverse(nums, 0, n-1)
		reverse(nums, k, n-1)
		reverse(nums, 0, k-1)
	} else {
		fmt.Println("not sure")
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	k := 2
	// direction := "right"
	direction := "left"

	fmt.Println(a)
	rotate(a, k, direction)
	fmt.Println(a)
}

// Approach 2
// Dry run example - Check Notion
func rotate(nums []int, k int) {
	var i, curr, count, n, next_pos, j int
	n = len(nums)
	k %= n
	for count < n {
		next_pos = (i + k) % n
		curr, nums[next_pos] = nums[next_pos], nums[i]
		count++
		j = next_pos
		for count < n && j != i {
			next_pos = (j + k) % n
			curr, nums[next_pos] = nums[next_pos], curr
			count++
			j = next_pos
		}
		i++
	}
}
