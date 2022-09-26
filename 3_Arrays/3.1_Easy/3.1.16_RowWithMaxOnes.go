package main

import (
	"fmt"
)

// Time complexity - O(n^2)
func RowWithMaxOnes(arr [][]int) int {
	global := 0
	var row int
	for i := 0; i < len(arr); i++ {
		local := 0
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 1 {
				local++
			}
		}
		if local > global {
			global = local
			row = i
		}
	}
	return row
}

// Time complexity - O(m+n)
// Start with right top corner
func RowWithOnes(arr [][]int) int {
	r, c := 0, len(arr)-1
	var res int
	for c >= 0 && r < len(arr) {
		if arr[r][c] == 1 {
			res = r
			c--
		} else {
			r++
		}
	}
	return res
}

func main() {
	arr := [][]int{{1, 1, 0}, {2, 3, 5}, {1, 1, 1}}
	fmt.Println("RowWithMaxOnes:", RowWithMaxOnes(arr))
	fmt.Println("RowWithMaxOnes:", RowWithOnes(arr))
}
