package leetcodeProject

// func main() {
// 	fmt.Println(numIslands([][]byte{
// 		{'1', '1', '1'},
// 		{'0', '1', '0'},
// 		{'1', '1', '1'},
// 	}))
// }

var way [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var mainMap [][]byte

func FindIslands(x int, y int) { //每次遍历，都会将从起始点发散开的所有1置为x
	for i := 0; i < 4; i++ {
		x += way[i][0]
		y += way[i][1]
		if (x >= 0 && x < len(mainMap)) && (y >= 0 && y < len(mainMap[0])) && mainMap[x][y] == '1' {
			mainMap[x][y] = 'x'
			FindIslands(x, y)
		}
		x -= way[i][0]
		y -= way[i][1]
	}
}
func NumIslands(grid [][]byte) int {
	var res int = 0
	mainMap = grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if mainMap[i][j] == '1' { //由于每次进入递归，都会将已经组成'孤岛'的1全部置为x，所以只要进入递归，说明这又是一片新的'孤岛'
				FindIslands(i, j)
				res++
			}
		}
	}

	return res
}
