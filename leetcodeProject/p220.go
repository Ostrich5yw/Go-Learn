package leetcodeProject

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(containsNearbyAlmostDuplicate([]int{-3, 3, -6}, 2, 3))
// }
/**
桶排序思想：
    以t + 1 作为桶基准来区分元素
    如数列 1 4 8 9 12 14 16         如果t为3 则t+1为4   bucket = n / (t+1)
    放入桶 0 1 2 2 3  3  4          桶只存放满足满足条件并且靠后的元素(索引2 索引3 同样在桶2,因为是向后遍历，所以存相对靠后的索引3元素即可)
*/
func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	numslen := len(nums)
	bucket := make(map[int]int, numslen)
	for i := 0; i < numslen; i++ {
		id := getId(nums[i], t)
		_, ok := bucket[id]
		if ok { //如果相同桶中有元素，直接返回
			return true
		}
		v, ok := bucket[id+1]
		if ok && abs(nums[i]-v) <= t { //如果是相邻桶中有元素，考虑有6/3=2 5/3=1这种情况  所以还需要具体比较两元素是否差值<=t
			return true
		}
		v, ok = bucket[id-1]
		if ok && abs(nums[i]-v) <= t {
			return true
		}
		bucket[id] = nums[i]
		if i >= k { //精髓：滑动窗口，一边向后遍历，一边删除当前位置之前k位，即不满足i-j <= k的元素
			delete(bucket, getId(nums[i-k], t))
		}
	}
	return false
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

func getId(n int, k int) int {
	if n >= 0 {
		return n / (k + 1)
	}
	return n/(k+1) - 1 //精髓：因为如果abs(n) < k ，比如-2 / 3 与 2 / 3 都会返回0，而负值-1则会将-2/3返回-1。由此将-2与2区分开来
} //同时因为所有负数统一-1，也不会影响其他负数。比如-4/3 将返回-2
