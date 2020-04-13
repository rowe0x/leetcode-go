package leetcode

func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
