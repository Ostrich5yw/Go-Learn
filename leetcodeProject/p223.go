package leetcodeProject

func maxAB(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minAB(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int { //左下右上，左下右上
	total := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	x1 := maxAB(ax1, bx1)	//两矩形找右侧边最小的，左侧边最大的
	x2 := minAB(ax2, bx2)
	y1 := maxAB(ay1, by1)	//上侧边最小的，下侧边最大的
	y2 := minAB(ay2, by2)
	if x1 > x2 || y1 > y2 {	//如果最小的右侧边小于最大的左侧边，或 ，最小的上侧边小于最大的下侧边，则无交集
		return total
	} else {
		return total - (y2-y1)*(x2-x1)
	}
}
