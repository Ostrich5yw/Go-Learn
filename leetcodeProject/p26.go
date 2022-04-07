package leetcodeProject

func RemoveDuplicates(nums []int) int {
	slow, quick := 1, 1
	for quick < len(nums) {
		if nums[quick] != nums[quick-1] {
			nums[slow] = nums[quick]
			slow++
		}
		quick++
	}
	return slow
}
