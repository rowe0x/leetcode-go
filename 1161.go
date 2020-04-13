package leetcode

func maxLevelSum(root *TreeNode) int {
	s := make(map[int]int)
	var dfs func(*TreeNode, int, map[int]int)
	dfs = func(node *TreeNode, level int, m map[int]int) {
		m[level] += node.Val
		if node.Left != nil {
			dfs(node.Left, level+1, m)
		}
		if node.Right != nil {
			dfs(node.Right, level+1, m)
		}
	}
	dfs(root, 1, s)
	var res = len(s)
	var max int
	for k, v := range s {
		if v > max {
			res = k
		}
		if v == max {
			if k < res {
				res = k
			}
		}
	}
	return res
}
