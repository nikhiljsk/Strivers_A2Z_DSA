package main

import (
	"fmt"
	"math"
)

func Divisors(n int) []int {
	var res []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(Divisors(36))
}
