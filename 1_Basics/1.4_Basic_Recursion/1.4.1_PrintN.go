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

func main() {
	printN(5)
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
