package leetcode

/*
func balancedString(s string) int {
	var ans int
	needBalanceCount := len(s)/4
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	for _, f := range freq {
		if f < needBalanceCount {
			ans += needBalanceCount-f
		}
	}
	return ans
}
 */
