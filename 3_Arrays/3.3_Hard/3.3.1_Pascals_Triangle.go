// Approach 1
func generate(n int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{1})
	if n == 1 {
		return res
	}

	for i := 1; i < n; i++ {
		res = append(res, generateNext(res[i-1]))
	}
	return res
}

func generateNext(arr []int) []int {
	res := make([]int, len(arr)+1)
	res[0] = 1
	ind := 1
	for i := 0; i < len(arr)-1; i++ {
		res[ind] = arr[i] + arr[i+1]
		ind++
	}
	res[ind] = 1
	return res
}

// Approach 2
func generate(n int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		res = append(res, make([]int, i+1))
		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}