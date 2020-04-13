package leetcode

func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	var tc int
	for n > 2 {
		tc = c
		c = a + b + c
		a, b = b, tc
		n--
	}
	return c

}
