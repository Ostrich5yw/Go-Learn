package leetcodeProject

import (
	"strconv"
)

func compareNum(nums []int, pos1 int, pos2 int) bool {
	num1 := strconv.Itoa(nums[pos1]) // 分别记录两个数的字符串，长度，当前位置
	num2 := strconv.Itoa(nums[pos2])
	len1 := int(float64(len(num1)))
	len2 := int(float64(len(num2)))
	mark1 := 0
	mark2 := 0

	for mark1 < len1 || mark2 < len2 {
		if mark1 >= len1 && mark2 < len2 { // 只要还未比较出两数大小，就一直循环遍历两数字符串
			mark1 = 0 // 比如8308 830   最终由于第一个数的'8'大于第二个数的'3'，所以返回false（即第五次循环）
		}
		if mark2 >= len2 && mark1 < len1 {
			mark2 = 0
		}

		if num1[mark1] < num2[mark2] {
			return true
		} else if num1[mark1] > num2[mark2] {
			return false
		} else {
			mark1++
			mark2++
			continue
		}
	}
	return true
}

func LargestNumber(nums []int) string {
	var res string = ""
	for i := 1; i < len(nums); i++ { // 冒泡排序，得到从大到小的nums序列（大小标准为compareNum）
		for j := i - 1; j >= 0; j-- {
			if compareNum(nums, j, j+1) {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
	for _, val := range nums { //数组拼接为字符串
		res += strconv.Itoa(val)
	}
	if res[0] == '0' { //考虑[0,0]情形
		res = "0"
	}
	return res
}
