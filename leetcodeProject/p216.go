package leetcodeProject

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(combinationSum3(9, 45))
// }

func recursion1(nums []int, res *[][]int, numslen int, nowN int, k int, n int) bool { //当前递归数组，结果数组，当前递归数组长度，当前递归数组和，目标长度，目标和
	if numslen == k && nowN == n { //如果长度，和 都满足条件，则加入结果数组res
		val := *res
		var t []int = make([]int, numslen) //注意，nums是指针，不能直接赋给res，必须复制数组元素到新数组
		copy(t, nums)
		addr := append(val, t)
		*res = addr //结果数组res是指针
		return true
	}
	if numslen >= k { //如果长度已经超出但是仍不满足条件，停止递归
		return false
	}
	for i := nums[numslen-1] + 1; i <= 9; i++ { //依次遍历当前最大值到9
		nums[numslen] = i //加入新元素并且长度+1
		numslen++
		recursion1(nums, res, numslen, nowN+i, k, n)
		numslen--
	}
	return true
}

func CombinationSum3(k int, n int) [][]int {
	var res [][]int
	for i := 1; i < 10; i++ {
		var nums []int = make([]int, k)
		nums[0] = i
		recursion1(nums, &res, 1, i, k, n)
	}
	return res
}
