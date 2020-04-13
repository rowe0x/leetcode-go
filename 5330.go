package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
	sum := make(map[*TreeNode]int)
	var sumDFS func(node *TreeNode) int
	sumDFS = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftVal := sumDFS(node.Left)
		rightVal := sumDFS(node.Right)
		sum[node] = leftVal + node.Val + rightVal
		return sum[node]
	}
	sumDFS(root)
	var ans = 0
	var productDFS func(node *TreeNode)
	productDFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			nodeLeftVal := sumDFS(node.Left)
			candidate := nodeLeftVal * sum[node.Left] % 1000000007
			if ans < candidate {
				ans = candidate
			}
			productDFS(node.Left)
		}
		if node.Right != nil {
			nodeRightVal := sumDFS(node.Right)
			candidate := nodeRightVal * sum[node.Right] % 100000007
			if ans < candidate {
				ans = candidate
			}
			productDFS(node.Right)
		}
	}
	productDFS(root)
	return ans
}