package leetcode

import (
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	m := height(root) + 1
	r := make([][]string, m)
	n := 1 << uint(m) - 1
	for i := 0; i < m; i++ {
		r[i] = make([]string, n)
	}
	fill(r, root, 0, 0, n)
	return r
}

func fill(res [][]string, root *TreeNode, level, lb, rb int) {
	if root == nil {
		return
	}
	mid := (lb+rb)/2
	res[level][mid] = strconv.Itoa(root.Val)
	fill(res, root.Left, level+1, lb, mid-1)
	fill(res, root.Right, level+1, mid+1,rb)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	lh := height(root.Left) + 1
	rh := height(root.Right) + 1
	if lh > rh {
		return lh
	}
	return rh
}
