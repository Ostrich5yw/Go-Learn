package leetcodeProject

import (
	d "DataStructure"
)

func NthUglyNumber(n int) int {
	mheap, mmap := d.NewHeap(), make(map[int]struct{})
	mheap.Push(1)
	pp := []int{2, 3, 5}
	for i := 1; ; i++ { //堆负责整体排序，map负责去除重复元素。堆顶每次存放当前丑数组最小元素，将其乘以2,3,5加入丑数组
		minV := mheap.Pop()
		if i == n {
			return minV
		}
		for _, v := range pp {
			temp := minV * v
			if _, ok := mmap[temp]; !ok {
				mheap.Push(temp)
				mmap[temp] = struct{}{}
			}
		}
	}
}
