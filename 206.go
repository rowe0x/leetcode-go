package leetcode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	var curr = head
	for curr != nil {
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre
}

