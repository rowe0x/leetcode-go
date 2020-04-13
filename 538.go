package leetcode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	right := convertBST(root.Right)
	if right != nil {
		root.Val += right.Val
	}
	left := convertBST(root.Left)
	if left != nil {
		left.Val += root.Val
	}
	return root
}
