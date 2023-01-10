func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func merge(arr [][]int) [][]int {
	// Sort the array based on 0th column element
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] < arr[j][0] {
			return true
		}
		return false
	})

	res := make([][]int, 0)
	l, r := arr[0][0], arr[0][1]
	for i := 1; i < len(arr); i++ {
		if arr[i][0] <= r {
			r = max(arr[i][1], r)
		} else {
			res = append(res, []int{l, r})
			l, r = arr[i][0], arr[i][1]
		}
	}
	res = append(res, []int{l, r})
	return res
}