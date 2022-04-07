package leetcodeProject

func MajorityElement(nums []int) int {
	mostNum := nums[0]
	lenNum := 0
	for i := 0; i < len(nums); i++ {
		if lenNum == 0 {
			mostNum = nums[i]
		}
		if nums[i] == mostNum {
			lenNum++
		} else {
			lenNum--
		}
	}
	return mostNum
}
