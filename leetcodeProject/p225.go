package leetcodeProject

type MyStack struct {
	A    []int
	Size int
	B    []int
}

func Constructor4() MyStack {
	return MyStack{A: make([]int, 100), Size: 0, B: make([]int, 100)}
}

func (stack *MyStack) Push(x int) {
	stack.A[0] = x							//先用A队列接收
	for i := 0; i < stack.Size; i++ {		//再将B队列所有元素插入A队列
		stack.A[i+1] = stack.B[i]
	}
	stack.Size++
	copy(stack.B, stack.A)					//将A队列复制到B队列
}											//保障了后入元素在队列最前，即后入先出

func (stack *MyStack) Pop() int {
	res := stack.B[0]
	for i := 1; i < stack.Size; i++ {
		stack.B[i-1] = stack.B[i]
	}
	stack.Size--
	return res
}

func (stack *MyStack) Top() int {
	return stack.B[0]
}

func (stack *MyStack) Empty() bool {
	return stack.Size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
