package main

import (
	"fmt"
)

func minSwap(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func GCD(a, b int) int {
	a, b = minSwap(a, b)
	for a > 1 {
		temp := b % a
		if temp == 0 {
			return a
		}
		a, b = temp, a
	}
	return a
}

func GCDRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRecursive(b, a%b)
}

func LCM(a, b, gcd int) int {
	return (a * b) / gcd
}

func main() {
	a, b := 3, 5
	gcd := GCD(a, b)
	lcm := LCM(a, b, gcd)
	fmt.Println(lcm, gcd)
	fmt.Println(GCDRecursive(a, b))
}
