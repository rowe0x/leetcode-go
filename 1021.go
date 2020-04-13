package leetcode

func removeOuterParentheses(S string) string {
	var ans []byte
	var block []byte
	var parenthesesCount int
	for i := 0; i < len(S); i++ {
		block = append(block, S[i])
		if S[i] == ')' {
			parenthesesCount++
		} else {
			parenthesesCount--
			if parenthesesCount == 0 {
				ans = append(ans, block[1:len(block)-1]...)
			}
		}
	}
	return string(ans)
}
