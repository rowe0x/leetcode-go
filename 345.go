package leetcode

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	ans := []byte(s)
	lp, rp := false, false
	for l < r {
		if lp && rp {
			ans[l], ans[r] = ans[r], ans[l]
			lp, rp = false, false
		} else if !lp {
			switch s[l] {
			case 'a','e','i','o','u':
				lp = true
			default:
				l++
			}
		} else {
			switch s[r] {
			case 'a','e','i','o','u':
				rp = true
			default:
				r--
			}
		}
	}
	return string(ans)
}
