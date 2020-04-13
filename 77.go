package leetcode

func combine(n int, k int) [][]int {
	var ans [][]int
	cand := make([]int, k)
	for i := 0; i < k; i++ {
		cand[i] = i
	}
	ans = append(ans, cand[:])
	for loop := 0; loop < k; loop++ {
		for i := k; i < n-k; i++ {
			cand[loop] = i
			ans = append(ans, cand[:])
		}
	}
	return ans
}
