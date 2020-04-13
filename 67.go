package leetcode

func addBinary(a string, b string) string {
	var ans []byte
	la, lb := len(a), len(b)
	m := la
	max := lb
	maxS := b
	if lb < la {
		m = lb
		max = la
		maxS = a
	}
	carry := false
	for i := 0; i < m; i++ {
		ai := a[len(a)-1-i]	- '0'
		bi := b[len(b)-1-i] - '0'
		if ai == 0 && bi == 0 {
			if carry {
				ans = append(ans, '1')
				carry = false
			} else {
				ans = append(ans, '0')
			}
		} else if ai == 1 && bi == 1 {
			if carry {
				ans = append(ans, '1')
			} else {
				ans = append(ans, '0')
				carry = true
			}
		} else {
			if carry {
				ans = append(ans, '0')
			} else {
				ans = append(ans, '1')
			}
		}
	}
	for i := m; i < max; i++ {
		mi := maxS[len(maxS)-1-i] - '0'
		if mi == 1 {
			if carry {
				ans = append(ans, '0')
			} else {
				ans = append(ans, '1')
			}
		} else {
			if carry {
				ans = append(ans, '1')
				carry = false
			} else {
				ans = append(ans, '0')
			}
		}
	}
	if carry {
		ans = append(ans, '1')
	}
	for i := len(ans)/2-1; i >= 0; i-- {
		opp := len(ans)-1-i
		ans[i], ans[opp] = ans[opp], ans[i]
	}
	return string(ans)
}
