package leetcode

func strStr(haystack string, needle string) int {
	lps := make([]int, len(needle))
	for i := 0; i < len(needle); i++ {
		for j := 0; j < i; j++ {
			if needle[j] == needle[i-j-1] {
				lps[i]++
			}
		}
	}
	ans := -1
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return j-len(needle)
			}
			continue
		}
		if j == 0 {
			i++
			continue
		}
		j = lps[j-1]
	}
	return ans
}
