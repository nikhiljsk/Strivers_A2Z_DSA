package main

import (
	"fmt"
)

func ref(a *int) *int {
	*a += 1
	return a
}

func value(a int) int {
	a += 1
	return a
}

func valueArr(arr []int) {
	arr[0] += 100
}

func refArr(arr *[]int) {
	(*arr)[0] += 100
}

func main() {
	var a int
	fmt.Println("Before value:", a)
	value(a)
	fmt.Println("After value:", a)
	fmt.Println("Before ref:", a)
	ref(&a)
	fmt.Println("After ref:", a)

	// Pass by reference by default
	arr := []int{1, 2, 3}
	fmt.Println("Before value:", arr)
	valueArr(arr)
	fmt.Println("After value:", arr)

	fmt.Println("Before ref:", arr)
	refArr(&arr)
	fmt.Println("After value:", arr)
}
