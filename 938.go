package leetcode

func rangeSumBST(root *TreeNode, L int, R int) int {
	var s int
	if root == nil {
		return 0
	}
	if root.Val >= L && root.Val <= R {
		s += root.Val
	}
	s += rangeSumBST(root.Left, L, R)
	s += rangeSumBST(root.Right, L, R)
	return s
}
