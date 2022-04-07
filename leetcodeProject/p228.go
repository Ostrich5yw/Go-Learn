package leetcodeProject

import "strconv"

func SummaryRanges(nums []int) []string {
	len := len(nums)
	res := make([]string, len)
	lenres := 0
	start := 0
	for i := 0; i < len; i++ {
		if i != len-1 && nums[i] == (nums[i+1]-1) {  // 第一个判断主要是为了区分最后一个元素，因为其没有对应的nums[i+1]
			continue
		} else {
			if start != i {		//start记录每一个字符串的起始
				res[lenres] = strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[i])
			} else {
				res[lenres] = strconv.Itoa(nums[i])
			}
			lenres++
			start = i + 1
		}
	}

	return res[0:lenres]
}
