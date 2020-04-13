package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	head := l
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l = l.Next
			l1 = l1.Next
		} else {
			l.Next = l2
			l = l.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	return head
}
