package leetcodeProject

func SearchMatrix(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] > target { //如果当前位置i,j大于target，说明之后i+1...行
				column = j //也不会超过j的位置
				break
			}
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
