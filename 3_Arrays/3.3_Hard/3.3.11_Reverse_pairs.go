// TLE
func reversePairs(nums []int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				c++
			}
		}
	}
	return c
}

// Approach 2
func merge(arr []int, left, mid, right int) int {
	res := 0
	i, j := left, mid+1
	for i <= mid {
		for j <= right && arr[i] > (2*arr[j]) {
			j++
		}
		res += (j - (mid + 1))
		i++
	}

	temp := make([]int, right+1)
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	if i > mid {
		copy(temp[k:], arr[j:right+1])
	} else if j > right {
		copy(temp[k:], arr[i:mid+1])
	}
	for i = left; i <= right; i++ {
		arr[i] = temp[i]
	}
	return res
}

func mergeSort(nums []int, left, right int) int {
	res := 0
	if left < right {
		mid := (left + right) / 2
		res += mergeSort(nums, left, mid)
		res += mergeSort(nums, mid+1, right)
		res += merge(nums, left, mid, right)
	}
	return res
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}