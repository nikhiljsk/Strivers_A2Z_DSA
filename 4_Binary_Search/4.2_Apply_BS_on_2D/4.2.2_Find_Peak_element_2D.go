// Explaination: https://leetcode.com/problems/find-a-peak-element-ii/solutions/1338377/java-c-detailed-explanation/
// Tip - Since any peak element can be returned, BS works here. And also no adjacent cells are equal.
func findPeakGrid(mat [][]int) []int {
	startCol, endCol := 0, len(mat[0])-1

	for startCol <= endCol {
		midCol := (startCol + endCol) / 2
		maxRow := 0

		// Find the max element row in midCol
		for row := 0; row < len(mat); row++ {
			if mat[maxRow][midCol] <= mat[row][midCol] {
				maxRow = row
			}
		}

		var leftIsBig, rightIsBig bool
		if midCol-1 >= startCol && mat[maxRow][midCol-1] > mat[maxRow][midCol] {
			leftIsBig = true
		}
		if midCol+1 <= endCol && mat[maxRow][midCol+1] > mat[maxRow][midCol] {
			rightIsBig = true
		}

		if leftIsBig {
			endCol = midCol - 1
		} else if rightIsBig {
			startCol = midCol + 1
		} else {
			return []int{maxRow, midCol}
		}
	}
	return []int{}
}