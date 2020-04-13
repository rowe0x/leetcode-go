package leetcode

func removePalindromeSub(s string) int {
	count := len(s)
	if count == 0 {
		return 0
	}
	left, right := 0, count-1
	for left <= right {
		if s[left] != s[right] {
			return 2
		}
		left++
		right--
	}
	return 1
}
