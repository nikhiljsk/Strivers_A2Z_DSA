// Approach 1
// Transpose and Reverse rows
func rotate(matrix [][]int) {
	R, C := len(matrix), len(matrix[0])

	// Transpose
	for i := 0; i < R; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse rows
	for i := 0; i < R; i++ {
		for j := 0; j < C/2; j++ {
			matrix[i][j], matrix[i][R-j-1] = matrix[i][R-j-1], matrix[i][j]
		}
	}
}

// Approach 2
// Spiral replace
func rotate(matrix [][]int) {
	n := len(matrix)
	depth := n / 2
	t, b, l, r := 0, n-1, 0, n-1
	for i := 0; i < depth; i++ {
		jLen := n - (2 * i) - 1
		for j := 0; j < jLen; j++ {
			temp := matrix[t][l+j]
			matrix[t][l+j] = matrix[b-j][l]
			matrix[b-j][l] = matrix[b][r-j]
			matrix[b][r-j] = matrix[t+j][r]
			matrix[t+j][r] = temp
		}
		r--
		l++
		b--
		t++
	}
}