package leetcode

func longestSubsequence(arr []int, difference int) int {
	var ans int
	pre := make(map[int]int)
	var p int
	for _, n := range arr {
		p = n-difference
		if pf, ok := pre[p]; ok {
			pre[n] = pf+1
			if ans < pre[n] {
				ans = pre[n]
			}
		} else {
			pre[n] = 1
		}
	}
	return p
}
