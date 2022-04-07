package leetcodeProject

import (
	"container/heap"
	d "DataStructure"
)

type MedianFinder struct {
	MaxItem *d.MaxHeap		//声明一个大顶堆，一个小顶堆   大顶堆用来存放前一半比较小的数，小顶堆用来存放后一半比较大的数，即小顶堆堆顶大于大顶堆堆顶
	MinItem *d.MinHeap
}

func Constructor6() MedianFinder {
	return MedianFinder{MaxItem: d.NewMaxHeap(), MinItem: d.NewMinHeap()}
}

/**
	先将新元素加入大顶堆，再将大顶堆堆顶元素（最大值）加入小顶堆
	始终保证大顶堆元素==小顶堆元素||小顶堆元素+1
	否则，将小顶堆堆顶（最小值）加入大顶堆

	比如 3,4,1,2
	3 3加入大堆 3加入小堆 “小堆多于大堆” 3加入大堆    小：   大：3
	4 4加入大堆 4加入小堆 “小堆等于大堆” 		     小：4  大：3
	1 1加入大堆 3加入小堆 “小堆多于大堆” 3加入大堆	  小：4  大：1 3
	2 2加入大堆 3加入小堆 “小堆等于大堆” 			 小：4 3 大：3
*/
func (h *MedianFinder) AddNum(num int) {
	heap.Push(h.MaxItem, num)
	heap.Push(h.MinItem, heap.Pop(h.MaxItem).(int))
	if h.MaxItem.Len() < h.MinItem.Len() {
		heap.Push(h.MaxItem, heap.Pop(h.MinItem).(int))
	}
}

func (h *MedianFinder) FindMedian() float64 {
	if h.MaxItem.Len() == h.MinItem.Len() {
		return float64(h.MaxItem.ReverseIntSlice[0]+h.MinItem.IntSlice[0]) / 2
	}
	return float64(h.MaxItem.ReverseIntSlice[0])
}
