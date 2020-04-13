package leetcode

import (
	"sort"
)

func heightChecker(heights []int) int {
	var res int
	dup := append([]int(nil), heights...)
	sort.Ints(dup)
	for i := range heights {
		if heights[i] != dup[i] {
			res++
		}
	}
	return res
}
