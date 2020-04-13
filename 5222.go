package leetcode

func balancedStringSplit(s string) int {
	Lcount, RCount := 0, 0
	var ans int
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			Lcount++
		} else {
			RCount++
		}
		if Lcount == RCount {
			ans++
			Lcount = 0
			RCount = 0
		}
	}
	return ans
}
