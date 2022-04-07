package leetcodeProject

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor9(matrix [][]int) NumMatrix { //sumMatrix[i][j]保存从(0,0)到(i,j)的元素和
	for i := 0; i < len(matrix); i++ {
		if i > 0 { //(0,0) = (0,0)
			matrix[i][0] += matrix[i-1][len(matrix[0])-1] //换行之后，(i,j)等于(i-1,len)+(i,j)
		}
		for j := 1; j < len(matrix[0]); j++ { //其他位置，(i,j)等于(i-1,j)+(i,j)
			matrix[i][j] += matrix[i][j-1]
		}
	}
	return NumMatrix{sumMatrix: matrix}
}

func (t *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		//考虑左上角(0,0)到右下角(i,j)的情况 以及 左上角(i,0)到右下角(i,j)
		if col1 == 0 {
			if i == 0 {
				sum += t.sumMatrix[i][col2] //原点直接加
				continue
			} //(i,j)等于 (i,j)-(i-1,len)
			sum += (t.sumMatrix[i][col2] - t.sumMatrix[i-1][len(t.sumMatrix[0])-1])
			continue
		}
		sum += (t.sumMatrix[i][col2] - t.sumMatrix[i][col1-1])
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
