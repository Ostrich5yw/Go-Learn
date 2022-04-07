package leetcodeProject

type MyQueue struct {
	A     []int
	B     []int
	sizeA int
	sizeB int
}

func Constructor5() MyQueue {
	return MyQueue{A: make([]int, 100), B: make([]int, 100), sizeA: 0, sizeB: 0}
}

func (tt *MyQueue) Push(x int) {
	tt.sizeB = 0
	for i := tt.sizeA - 1; i >= 0; i-- {
		tt.B[tt.sizeB] = tt.A[i]
		tt.sizeB++
	}
	tt.B[tt.sizeB] = x
	tt.sizeA = 0
	for i := tt.sizeB; i >= 0; i-- {
		tt.A[tt.sizeA] = tt.B[i]
		tt.sizeA++
	}
}

func (tt *MyQueue) Pop() int {
	tt.sizeA--
	return tt.A[tt.sizeA]
}

func (tt *MyQueue) Peek() int {
	return tt.A[tt.sizeA-1]
}

func (tt *MyQueue) Empty() bool {
	return tt.sizeA == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
