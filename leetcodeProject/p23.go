package leetcodeProject

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{1, nil}
	rec := res
	for list1 != nil || list2 != nil {
		if list2 == nil {
			rec.Next = list1
			break
		}
		if list1 == nil {
			rec.Next = list2
			break
		}
		if list1.Val < list2.Val {
			rec.Next = list1
			rec = rec.Next
			list1 = list1.Next
		} else {
			rec.Next = list2
			rec = rec.Next
			list2 = list2.Next
		}
	}
	return res.Next
}

func merge(lists []*ListNode) *ListNode { // 分治法
	Len := len(lists)
	if Len == 0 { // 针对输入为空的情况
		return nil
	}
	if Len == 1 {
		return lists[0]
	}
	l := merge(lists[:Len/2])
	r := merge(lists[Len/2:])
	return mergeTwoLists(l, r)
}

func MergeKLists(lists []*ListNode) *ListNode {

	return merge(lists)
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode, ch chan *ListNode) {
	res := &ListNode{1, nil}
	rec := res
	for list1 != nil || list2 != nil {
		if list2 == nil {
			rec.Next = list1
			break
		}
		if list1 == nil {
			rec.Next = list2
			break
		}
		if list1.Val < list2.Val {
			rec.Next = list1
			rec = rec.Next
			list1 = list1.Next
		} else {
			rec.Next = list2
			rec = rec.Next
			list2 = list2.Next
		}
	}
	ch <- res.Next
}

func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ch1 := make(chan *ListNode, len(lists))
	for _, v := range lists {
		ch1 <- v
	}
	Len := len(lists)
	for Len != 1 {
		target := Len / 2
		if Len%2 != 0 {
			target++
		}
		for i := 0; i < Len/2; i++ {
			go mergeTwoLists2(<-ch1, <-ch1, ch1)
		}
		for target != len(ch1) {
		}
		Len = target
	}
	return <-ch1
}
