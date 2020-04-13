package leetcode

func multiply(num1 string, num2 string) string {
	var n1 int64
	for i := 0; i < len(num1); i++ {
		n1 = 10 * n1 + int64(num1[i] - '0')
	}
	var n2 int64
	for i := 0; i < len(num2); i++ {
		n2 = 10 * n2 + int64(num2[i] - '0')
	}
	n := n1*n2
	var ans []byte
	for n > 0 {
		ans = append(ans, byte(n%10)+'0')
		n /= 10
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return string(ans)
}
