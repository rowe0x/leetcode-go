package leetcode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var ans [][]int
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, s int, nodes []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == s {
				ans = append(ans, append(nodes, node.Val))
			}
			return
		}
		dfs(node.Left, s-node.Val, append(nodes, node.Val))
		dfs(node.Right, s-node.Val, append(nodes, node.Val))
	}
	dfs(root, sum, []int{})
	return ans
}
