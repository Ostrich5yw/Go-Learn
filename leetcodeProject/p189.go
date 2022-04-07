package leetcodeProject

import "fmt"

/**
对于1,2,3,4,5,6 后移3位 从索引0开始
0到3,3到0 此时0重复，向后一位，由1开始
1到4,4到1 此时1重复，向后一位，由2开始
2到5,5到2	完成

有些可以直接走完：
1,2,3,4,5  后移2位 从索引0开始
0到2,2到4,4到1,1到3  完成
**/
func Rotate(nums []int, k int) {
	numslen := len(nums)
	var des int
	var recLast int
	recMap := make(map[int]bool)
	for i := 0; i < numslen; i++ {
		if len(recMap) == numslen { //map中有所有元素，说明循环完成
			break
		}
		recLast = nums[i]       //记录每轮起始元素
		des = (i + k) % numslen //记录起始元素后移k位的最终位置
		for !recMap[des] {      // 如果这个数已经见过，跳出本轮循环。序数+1后继续循环
			recMap[des] = true
			temp := recLast
			recLast = nums[des] //1,2,3,4,5   若1向后移2格。则recLast记录1，des指向3 -> recLast记录3，nums[des]变为1，des指向5
			nums[des] = temp
			des = (des + k) % numslen
		}
	}
	// nums[1] = recLast
	fmt.Println(nums)
}
