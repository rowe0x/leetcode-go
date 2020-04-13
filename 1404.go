package leetcode

func numSteps(ss string) int {
	s := []byte(ss)
	var ans int
	var carry bool
	right := len(s) - 1
	for right > 0 || carry {
		if !carry {
			if s[right] == '0' {
				right--
			} else {
				carry = true
			}
		} else {
			if right == -1 {
				break
			}
			if s[right] == '0' {
				carry = false
			}
			right--
		}
		ans++
	}
	return ans
}
