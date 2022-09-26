package main

import (
	"fmt"
)

func printN(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	printN(n - 1)
}

func backtrackPrintN(i, n int) {
	if i > n {
		return
	}
	backtrackPrintN(i+1, n)
	fmt.Println(i)
}

func main() {
	printN(5)
	fmt.Println("Backtracking results:")
	backtrackPrintN(1, 5)
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
