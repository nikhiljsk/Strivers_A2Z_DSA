package main

import (
	"fmt"
)

func Merge(arr []int, l, mid, r int) {
	// Two arrays, one starts with index i, other with mid+1
	i := l
	j := mid + 1
	k := l
	temp := make([]int, r+1) // New array to store result
	// Copy whatever element is smallest
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	// Copy the remaining incase we couldn't iterate
	if i > mid {
		copy(temp[k:], arr[j:r+1])
	} else if j > r {
		copy(temp[k:], arr[i:mid+1])
	}
	// Transfer the result to the array for inplace
	for i = l; i <= r; i++ {
		arr[i] = temp[i]
	}
}

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	Merge(arr, l, mid, r)
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 5}
	fmt.Println("Before Sorting:", arr)
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println("After Sorting:", arr)
}
