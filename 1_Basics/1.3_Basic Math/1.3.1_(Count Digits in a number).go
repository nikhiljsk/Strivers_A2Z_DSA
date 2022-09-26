// Given a number N. Count the number of digits in N which evenly divides N.
package main

import (
	"fmt"
)

func isDivisible(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false
}

func getDigits(n int) []int {
	var res []int
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	return res
}

func countDigits(n int) int {
	digits := getDigits(n)
	var res int
	for _, digit := range digits {
		if isDivisible(n, digit) {
			res++
		}
	}
	return res
}

func main() {
	n := 23
	// n=8394
	fmt.Println(countDigits(n))
}
