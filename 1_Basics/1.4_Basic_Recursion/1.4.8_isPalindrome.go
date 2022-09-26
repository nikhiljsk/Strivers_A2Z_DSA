package main

import (
	"fmt"
)

func isPalindrome(arr []byte, left, right int) bool {
	if left >= right {
		return true
	}
	if arr[left] != arr[right] {
		return false
	}
	return isPalindrome(arr, left+1, right-1)
}

func main() {
	s := "abbba"
	fmt.Println("isPalindrome:", isPalindrome([]byte(s), 0, len(s)-1))
}

// Time Complexity
// Time - O(n)
// Storage - O(n)
