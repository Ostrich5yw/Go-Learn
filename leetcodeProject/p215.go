package leetcodeProject

func FindKthLargest(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		j := i
		temp := nums[j]
		for ; j > 0; j-- {
			if temp < nums[j-1] {
				nums[j] = nums[j-1]
				continue
			}
			break
		}
		nums[j] = temp
	}
	return nums[len(nums)-k]
}
