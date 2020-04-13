package leetcode

func balancedString(s string) int {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	var l, r int
	ans := len(s)
	max := len(s)/4
	for r < len(s) {
		if freq['Q'] > max || freq['W'] > max || freq['E'] > max || freq['R'] > max {
			r++
			freq[s[r]]--
		} else {
			if ans < r-l {
				ans = r-l
			}
			l++
			freq[s[l]]++
		}
	}
	return ans
}
