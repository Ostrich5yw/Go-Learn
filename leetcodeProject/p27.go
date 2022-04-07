package leetcodeProject

func RemoveElement(nums []int, val int) int {
	slow, quick := 0, 0
	for ; quick < len(nums); quick++ {
		if nums[quick] != val {
			nums[slow] = nums[quick]
			slow++
		}
	}
	return slow
}
