package main

import (
	"fmt"
)

// The problem has various question follow -ups
// Check both GFG and Leetcode. All the three methods will be useful

// Approach 1
// Find the corresponding row in log(m)
// Find the target in row in log(n)
// O(log(m)+log(n)) == O(log(m*n))
func findRow(arr [][]int, n int) int {
	t, b := 0, len(arr)-1
	for t <= b {
		mid := t + ((b - t) / 2)
		if arr[mid][len(arr)-1] >= n {
			b = mid - 1
		} else {
			t = mid + 1
		}
	}
	if t >= len(arr) {
		return 0
	}
	return t
}

func searchElement(arr [][]int, row, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + ((r - l) / 2)
		val := arr[row][mid]
		if val == target {
			return true
		} else if val >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

// Approach 2
// Start from right top corner and navigate accordingly
// Optimal approach for GFG problem
// O(m+n)
func searchNavigate(arr [][]int, target int) bool {
	r, c := 0, len(arr)-1
	for r < len(arr) && c >= 0 {
		val := arr[r][c]
		if val == target {
			return true
		}
		if val > target {
			c--
		} else {
			r++
		}
	}
	return false
}

// Approach 3
// Optimal approach for Leetcode problem
// You can assume all the elements to be in a single flattened array by using this:
// col = [mid%c], row = [mid/c], where c = number of columns
// O(log(m*n))
func searchFlatArray(arr [][]int, target int) bool {
	l, r, c := 0, len(arr)*len(arr[0])-1, len(arr[0])
	for l <= r {
		mid := l + ((r - l) / 2)
		if arr[mid/c][mid%c] == target {
			return true
		}
		if arr[mid/c][mid%c] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	n := 8
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	row := findRow(arr, n)
	fmt.Println("Approach 1:", searchElement(arr, row, n))
	fmt.Println("Approach 2:", searchNavigate(arr, n))
	fmt.Println("Approach 3:", searchFlatArray(arr, n))
}
