package leetcode

import (
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	var res = 1
	var curIdx = 0
	for i := 0; i < len(pairs)-1; i++ {
		if pairs[curIdx][1] < pairs[i+1][0] {
			res++
			curIdx = i+1
		}
	}
	return res
}
