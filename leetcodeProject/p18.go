package leetcodeProject

import (
	"math"
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	res := [][]int{}
	nLen := len(nums)
	sort.Ints(nums)
	iBef, jBef := math.MaxInt, math.MaxInt
	for i := 0; i < nLen; i++ {
		if iBef == nums[i] {
			continue
		}
		for j := i + 1; j < nLen; j++ {
			if jBef == nums[j] {
				continue
			}
			left, right := j+1, nLen-1
			for left < right {
				total := nums[i] + nums[j] + nums[left] + nums[right]
				if total == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left, right = left+1, right-1
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right++
					}
				} else if total > target {
					right--
				} else {
					left++
				}
			}
			jBef = nums[j]
		}
		iBef = nums[i]
		jBef = math.MaxInt
	}
	return res
}
