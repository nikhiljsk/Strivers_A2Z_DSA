package main

import (
	"fmt"
)

func sumN(n, sum int) int {
	if n <= 0 {
		return sum
	}
	return sumN(n-1, sum+n)
}

func main() {
	fmt.Println("Sum:", sumN(10, 0))
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
