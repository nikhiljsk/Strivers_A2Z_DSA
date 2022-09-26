package main

import (
	"fmt"
)

func printN(n, i int) {
	if i > n {
		return
	}
	fmt.Println(i)
	printN(n, i+1)
}

func backtrackPrintN(n int) {
	if n <= 0 {
		return
	}
	backtrackPrintN(n - 1)
	fmt.Println(n)
}

func main() {
	printN(5, 0)
	fmt.Println("Backtracking results:")
	backtrackPrintN(5)
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
