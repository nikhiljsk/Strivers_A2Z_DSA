package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 5, 10, 15, 10, 5}
	freq := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if _, ok := freq[arr[i]]; ok {
			freq[arr[i]] += 1
		} else {
			freq[arr[i]] = 1
		}
	}
	fmt.Println("Number Frequencies:")
	for k, v := range freq {
		fmt.Printf("%v:%v\n", k, v)
	}
}
