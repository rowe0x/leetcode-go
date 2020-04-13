package leetcode
import "sort"

func arrayRankTransform(arr []int) []int {
	t := make([]int, len(arr))
	copy(t, arr)
	sort.Sort(sort.IntSlice(t))
	rank := make(map[int]int)
	curRank := 1
	for _, n := range t {
		if _, ok := rank[n]; !ok {
			rank[n] = curRank
			curRank++
		}
	}
	var ans []int
	for _, n := range arr {
		ans = append(ans, rank[n])
	}
	return ans
}
