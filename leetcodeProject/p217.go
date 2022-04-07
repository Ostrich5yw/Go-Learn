package leetcodeProject

import "fmt"

func ContainsDuplicate(nums []int) bool {
	numslen := len(nums)
	hashmap := make(map[int]bool, numslen)
	for i := 0; i < numslen; i++ {
		_, ok := hashmap[nums[i]]
		fmt.Println(ok)
		if ok {
			return true
		} else {
			hashmap[nums[i]] = true
		}
	}
	return false
}
