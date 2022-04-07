package leetcodeProject

func FindDuplicate(nums []int) int {
	slow, quick := 0, 0
	for {
		slow = nums[slow]
		quick = nums[nums[quick]]
		if nums[slow] == nums[quick] { //两指针相交， 一指针回退到起点，再次相交时则为环入口
			quick = 0
			for nums[slow] != nums[quick] {
				quick = nums[quick]
				slow = nums[slow]
			}
			return nums[slow]
		}
	}
}
