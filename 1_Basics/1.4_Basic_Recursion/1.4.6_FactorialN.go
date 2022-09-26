package main

import (
	"fmt"
)

func factN(n, res int) int {
	if n <= 0 {
		return res
	}
	return factN(n-1, res*n)
}

func main() {
	fmt.Println("Fact:", factN(4, 1))
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
