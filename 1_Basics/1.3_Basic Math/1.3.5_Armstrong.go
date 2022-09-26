package main

import (
	"fmt"
)

func getDigits(n int) []int {
	var res []int
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	return res
}

func isArmstrong(n int) bool {
	digits := getDigits(n)
	var sum int
	for _, digit := range digits {
		sum += (digit * digit * digit)
	}
	return sum == n
}

func main() {
	fmt.Println(isArmstrong(153))
}
