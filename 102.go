package leetcode

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) < (depth+1) {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
		return
	}
	dfs(root, 0)
	return res
}
