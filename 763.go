package leetcode

func partitionLabels(S string) []int {
	maxIdx := make([]int, 26)
	for i := 0; i < len(S); i++ {
		maxIdx[S[i]-'a'] = i
	}
	var ans []int
	var curMaxIdx = 0
	var curCount = 0
	for i := 0; i < len(S); i++ {
		if curMaxIdx < maxIdx[S[i]-'a'] {
			curMaxIdx = maxIdx[S[i]-'a']
		}
		curCount++
		if i == curMaxIdx {
			ans = append(ans, curCount)
			curCount = 0
		}
	}
	return ans
}
