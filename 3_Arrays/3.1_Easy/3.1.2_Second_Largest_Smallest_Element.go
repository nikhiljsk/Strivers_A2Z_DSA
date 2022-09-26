package main

import (
	"fmt"
)

func secondSmall(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	s1, s2 := 2147483647, 2147483647
	for _, v := range arr {
		if v < s1 {
			s2 = s1
			s1 = v
		} else if v < s2 && s1 != v {
			s2 = v
		}
	}
	return s2
}

func secondLarge(arr []int) int {
	l1, l2 := -2147483648, -2147483648
	for _, v := range arr {
		if v > l1 {
			l2 = l1
			l1 = v
		} else if v > l2 && l1 != v {
			l2 = v
		}
	}
	return l2
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Second Small:", secondSmall(arr))
	fmt.Println("Second Large:", secondLarge(arr))
}
