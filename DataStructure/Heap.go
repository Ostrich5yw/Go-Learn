package DataStructure

import (
	"sort"
)

//真小顶堆
type MinHeap struct {
	sort.IntSlice
}

func NewMinHeap() *MinHeap {
	return &MinHeap{IntSlice: sort.IntSlice{}}
}

func (h *MinHeap) Push(a interface{}) {
	h.IntSlice = append(h.IntSlice, a.(int))
}

func (h *MinHeap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type ReverseIntSlice []int

func (x ReverseIntSlice) Len() int           { return len(x) }
func (x ReverseIntSlice) Less(i, j int) bool { return x[i] > x[j] }
func (x ReverseIntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type MaxHeap struct {
	ReverseIntSlice
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{ReverseIntSlice: ReverseIntSlice{}}
}

func (h *MaxHeap) Push(a interface{}) {
	h.ReverseIntSlice = append(h.ReverseIntSlice, a.(int))
}

func (h *MaxHeap) Pop() interface{} {
	a := h.ReverseIntSlice
	v := a[len(a)-1]
	h.ReverseIntSlice = a[:len(a)-1]
	return v
}

//这只是一个用sort方法仿制的堆，耗时仍然等于对一个数组进行排序
type Heap struct {
	HeapItem sort.IntSlice
}

func NewHeap() (h *Heap) {
	return &Heap{HeapItem: sort.IntSlice{}}
}

func (h *Heap) Push(a int) {
	h.HeapItem = append(h.HeapItem, a)
	sort.Sort(h.HeapItem)
}

func (h *Heap) Pop() int {
	temp := h.HeapItem[0]
	h.HeapItem = h.HeapItem[1:]
	return temp
}
