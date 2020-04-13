package leetcode

func myAtoi(str string) int {
	var negative bool
	var ans int
	var first = true
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ' ':
			if first {
				continue
			}
			goto RET
		case '-':
			if first {
				first = false
				negative = true
			} else {
				goto RET
			}
		case '+':
			if first {
				first = false
			} else {
				goto RET
			}
		case '0','1','2','3','4','5','6','7','8','9':
			ans = 10 * ans + int(str[i]-'0')
			if ans > 2147483648 {
				goto RET
			}
		default:
			goto RET
		}
	}

RET:
	if negative {
		if ans > 2147483648 {
			return -2147483648
		}
		return -ans
	} else if ans >= 2147483648 {
		return 2147483647
	}
	return ans
}
