package leetcodeProject

import "math"

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(rob([]int{200, 3, 140, 20, 10}))
// }

func findMax(nums []int) int {
	max := -1
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func Rob(nums []int) int {
	max := -1
	if len(nums) <= 3 { //数组长度小于3，返回最大值
		return findMax(nums)
	}
	nums1 := make([]int, len(nums)) //一共进行两次不相关的遍历，分别从0-n-1与1-n，比较得到两数组中的最大值，就是返回结果
	copy(nums1, nums[:])
	nums1[1] = int(math.Max(float64(nums1[0]), float64(nums1[1])))
	for i := 2; i < len(nums)-1; i++ {
		nums1[i] = int(math.Max(float64(nums1[i-1]), float64(nums1[i-2]+nums1[i])))
		max = int(math.Max(float64(nums1[i]), float64(max)))
	}
	nums[2] = int(math.Max(float64(nums[1]), float64(nums[2])))
	for i := 3; i < len(nums); i++ {
		nums[i] = int(math.Max(float64(nums[i-1]), float64(nums[i-2]+nums[i])))
		max = int(math.Max(float64(nums[i]), float64(max)))
	}
	return max
}
