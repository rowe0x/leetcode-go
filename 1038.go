package leetcode

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Right != nil {
		root.Val += bstToGst(root.Right).Val
	}
	if root.Left != nil {
		root.Left.Val += root.Val
	}
	return root
}