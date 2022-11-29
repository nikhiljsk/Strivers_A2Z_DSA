func spiralOrder(matrix [][]int) []int {
	startRow, endRow := 0, len(matrix)-1
	startCol, endCol := 0, len(matrix[0])-1
	dir := 0
	var res []int

	for startRow <= endRow && startCol <= endCol {
		switch dir {
		// Left to right
		case 0:
			for i := startCol; i <= endCol; i++ {
				res = append(res, matrix[startRow][i])
			}
			startRow++
		// Top to down
		case 1:
			for i := startRow; i <= endRow; i++ {
				res = append(res, matrix[i][endCol])
			}
			endCol--
		// Right to left
		case 2:
			for i := endCol; i >= startCol; i-- {
				res = append(res, matrix[endRow][i])
			}
			endRow--
		// Down to top
		case 3:
			for i := endRow; i >= startRow; i-- {
				res = append(res, matrix[i][startCol])
			}
			startCol++
		}
		dir = (dir + 1) % 4
	}
	return res
}