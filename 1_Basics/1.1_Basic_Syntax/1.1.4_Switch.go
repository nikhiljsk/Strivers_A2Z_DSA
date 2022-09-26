package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)
	switch {
	case n%2 == 0:
		fmt.Println("This is even", n)
	case n%2 != 0:
		fmt.Println("This is odd", n)
		fallthrough
	default:
		fmt.Println("Fell from odd")
	}
}
