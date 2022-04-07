package leetcodeProject

import "sort"

func HIndex(citations []int) int {
	mcitations, mlen := sort.IntSlice(citations), len(citations)
	sort.Sort(mcitations)
	if (mlen == 1 && citations[0] == 0) || (mcitations[mlen-1] == 0) { //排出[0]和[0,0,0]情况
		return 0
	}
	for i := 0; i < mlen; i++ {
		oth := mlen - i
		if oth <= mcitations[i] { //0,1,3,5,6   遍历到3，如果大于3的个数，小于mcitations[i]本身，则返回个数
			return oth
		}
	}
	return mlen
}
