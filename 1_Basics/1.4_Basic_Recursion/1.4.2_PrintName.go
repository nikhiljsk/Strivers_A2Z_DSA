package main

import (
	"fmt"
)

func printName(n int, name string) {
	if n <= 0 {
		return
	}
	fmt.Println(name)
	printName(n-1, name)
}

func main() {
	printName(10, "Nikhil JSK")
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
