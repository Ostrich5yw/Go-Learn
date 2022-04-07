package leetcodeProject

func RemoveElements(head *ListNode, val int) *ListNode {
	var newhead *ListNode = &ListNode{Val: 0, Next: head}
	var point *ListNode = newhead
	for point != nil && point.Next != nil {
		if point.Next.Val == val {
			point.Next = point.Next.Next
			continue //注意这个continue，如果发生删除，无需进行移位操作
		}
		point = point.Next
	}
	return newhead.Next
}
