package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, bool)
	dfs = func(node *TreeNode) (*TreeNode, bool) {
		if node == nil {
			return nil, false
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == target {
				return nil, true
			}
			return node, false
		}
		lr, rr := false, false
		node.Left, lr = dfs(node.Left)
		node.Right, rr = dfs(node.Right)
		if lr || rr {
			return dfs(node)
		}
		return node, false
	}
	node, _ := dfs(root)
	return node
}
