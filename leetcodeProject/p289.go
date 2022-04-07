package leetcodeProject

var way1 [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} //控制搜索一个位置相邻的8个位置

func dfs3(board [][]int, x, y int) {
	oneNum := 0
	xLen, yLen := len(board), len(board[0])
	if y >= yLen { //搜索从左向右 从上至下依次遍历，当走到最右边，则向下一行同时列置为0，如果已经是最后一行，说明遍历完成，则return
		if x < xLen-1 {
			x, y = x+1, 0
		} else {
			return
		}
	}
	for i := 0; i < len(way); i++ { //依次查看该位置的8个相邻点
		tempx := x + way1[i][0]
		tempy := y + way1[i][1]
		if tempx >= xLen || tempx < 0 || tempy >= yLen || tempy < 0 { //防止越界
			continue
		}
		// fmt.Println(tempx, tempy, board[tempx][tempy])
		/**
		由于题目同时强调所有位置要同一时间变换，即在原矩阵下判断周围为1和0的个数；同时也强调要原地变换，即不引入新矩阵下，记录变换后新的值
		所以引入2,3  分别表示为原为0，后为1  和   原为1，后为0
		*/
		switch board[tempx][tempy] { //  2 代表该位置 0——>1  3 代表该位置 1——>0 		oneNum记录周围为1的点数
		case 1: //引入2,3主要是为了原地变换，既可以记录原来该位置的值，又可以在return时将其变为转换后的值
			oneNum++
		case 2:
			continue
		case 0:
			continue
		case 3:
			oneNum++
		}
	}
	// fmt.Println(oneNum)
	if board[x][y] == 1 { //按照规则与oneNum数量进行变换
		if oneNum < 2 {
			board[x][y] = 3
		} else if oneNum <= 3 {
			board[x][y] = 1
		} else {
			board[x][y] = 3
		}
	} else {
		if oneNum == 3 {
			board[x][y] = 2
		} else {
			board[x][y] = 0
		}
	}
	dfs3(board, x, y+1)
	if board[x][y] == 3 { //遍历完成，将3置为0 2置为1
		board[x][y] = 0
	} else if board[x][y] == 2 {
		board[x][y] = 1
	}
}
func GameOfLife(board [][]int) [][]int {
	dfs3(board, 0, 0)

	return board
}
