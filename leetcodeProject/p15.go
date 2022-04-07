package leetcodeProject

import (
	"math"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var res [][]int
	befi := math.MaxInt
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if befi == nums[i] { //为i去重
			continue
		}
		target, left, right := 0-nums[i], i+1, len(nums)-1 //双指针法，由于数组已经排序，且已经有目标值0-nums[i]
		for left < right {
			if nums[left]+nums[right] > target { //左右指针分别在数组两端，大于target则右指针左移，小于target则左指针右移
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				temp := []int{nums[i], nums[left], nums[right]}
				res = append(res, temp)
				left++
				right--
				for nums[left] == nums[left-1] && left < right { //为left去重
					left++
				}
				for nums[right] == nums[right+1] && left < right { //为right去重
					right--
				}
			}
		}
		befi = nums[i]
	}
	return res
}
