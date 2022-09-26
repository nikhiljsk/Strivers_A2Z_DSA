package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	s1 := fibonacci(n - 1)
	s2 := fibonacci(n - 2)
	return s1 + s2
}

func main() {
	n := 6
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

// Time Complexity
// Time - O(2^n)
// Storage - O(n)
