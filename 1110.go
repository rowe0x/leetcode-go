package leetcode

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	dn := make(map[int]struct{})
	for _, n := range to_delete {
		dn[n] = struct{}{}
	}
	var res []*TreeNode
	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(root *TreeNode, isRoot bool) *TreeNode {
		if root == nil {
			return nil
		}
		var rootDeleted bool
		if _, ok := dn[root.Val]; ok {
			rootDeleted = true
		}
		if isRoot && !rootDeleted {
			res = append(res, root)
		}
		root.Left = dfs(root.Left, rootDeleted)
		root.Right = dfs(root.Right, rootDeleted)
		if rootDeleted {
			return nil
		}
		return root
	}
	dfs(root, true)
	return res
}

