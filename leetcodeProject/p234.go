package leetcodeProject

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	num := 0
	nums := make([]int, 10*10*10*10*5) //nums记录链表前一半的元素
	quick, slow := head, head
	for quick.Next != nil && quick.Next.Next != nil { //快指针每次后移2个，慢指针每次后移一个，快指针遍历完成，慢指针刚好走到一半
		nums[num] = slow.Val
		num++
		slow = slow.Next
		quick = quick.Next.Next
	}
	if quick.Next != nil { //考虑链表元素个数为偶数，nums需要再记录一个
		nums[num] = slow.Val
		num++
	}
	slow = slow.Next                     //slow指向后一半元素的第一个
	for num = num - 1; num >= 0; num-- { //依次比较是否相等
		if nums[num] != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return true
}
