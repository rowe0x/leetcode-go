package leetcode

func removeZeroSumSublists(head *ListNode) *ListNode {
	var nodes []*ListNode
	var newHead = head
	var needAdjust bool
	node := head
	sum := 0
	pre := make(map[int]struct{})
	for node != nil {
		if needAdjust {
			if len(nodes) > 0 {
				nodes[len(nodes)-1].Next = node
			} else {
				newHead = node
			}
		}
		nodes = append(nodes, node)
		sum += node.Val
		if sum == 0 {
			pre = map[int]struct{}{}
			needAdjust = true
			newHead = nil
		} else {
			if _, ok := pre[sum]; ok {
				preSum := sum - nodes[len(nodes)-1].Val
				nodes = nodes[0:len(nodes)-1]
				for preSum != sum {
					delete(pre, preSum)
					preSum -= nodes[len(nodes)-1].Val
					nodes = nodes[0:len(nodes)-1]
				}
				needAdjust = true
			} else {
				pre[sum] = struct{}{}
			}
		}
		node = node.Next
	}
	if len(nodes) > 1 {
		nodes[len(nodes)-1].Next = nil
	} else {
		newHead = nil
	}
	return newHead
}