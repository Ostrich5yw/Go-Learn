package leetcodeProject

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	/**
	*	1->2->3->4->5
	*   newhead 指向1
	*   fp 指向2
	*	rec 指向3
	*   fp newhead交换方向，之后newhead指向新的头结点2，fp向后一位指向3，rec继续向后一位
	**/
	newhead := head
	fp := head.Next
	rec := fp.Next
	newhead.Next = nil
	for rec != nil {
		fp.Next = newhead //交换方向
		//三个指针整体后移一位
		newhead = fp
		fp = rec
		rec = rec.Next
	}
	fp.Next = newhead //因为rec为nil时fp刚好指向最后一位，所以还要进行一次交换
	newhead = fp
	return newhead
}
