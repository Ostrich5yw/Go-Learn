package leetcodeProject

func HIndex2(citations []int) int {
	mlen, res := len(citations), 0
	if (mlen == 1 && citations[0] == 0) || (citations[mlen-1] == 0) { //排出[0]和[0,0,0]情况
		return 0
	}
	for left, right := 0, mlen-1; left <= right; {
		mid := (left + right) / 2
		oth := mlen - mid
		if oth <= citations[mid] { //找到满足条件的了，应该往左侧继续寻找，因为 4,4,6,6,8,9,10  这种搜到6后，他左侧的6，oth更大，达到5
			res = oth
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
