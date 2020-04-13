package leetcode

func numOfWays(n int) int {
	a, b, mod := 6, 6, 1000000007
	for i := 0; i < n; i++ {
		m, n := a, b
		a = (3*m + 2*n) % mod
		b = (2*m + 2*n) % mod
	}
	return (a + b) % mod
}
