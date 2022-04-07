package leetcodeProject

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
// }

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numslen := len(nums)
	mindis := 10000 //记录两个任意的相等元素，存在的最近距离
	hashmap := make(map[int]int, numslen)
	for i := 0; i < numslen; i++ {
		v, ok := hashmap[nums[i]]
		hashmap[nums[i]] = i    //hashmap记录每一个元素最近一次出现的位置
		if ok && i-v < mindis { //如果当前位置和最近一次出现的位置，相距小于mindis，则mindis更新
			mindis = i - v
			if mindis < k { //小于要求，直接返回
				return true
			}
		}
	}
	return false
}
