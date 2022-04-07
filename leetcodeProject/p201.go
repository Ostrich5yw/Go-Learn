package leetcodeProject

func RangeBitwiseAnd(left int, right int) int { //规律：left —— right所有数字相与 == left 与 right的公共前缀，并将前缀后续补0
	var res int = 0
	for left != right {
		left = left >> 1
		right = right >> 1
		res++
	}
	return left << res
}
