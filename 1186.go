package leetcode

func maximumSum(arr []int) int {
	res := -10000
	s := 0
	pre := 0
	validBlock := false
	var t int
	for _, n := range arr {
		s += n
		if s < 0 {
			if validBlock {
				t = s - n + pre
				pre = s - n
			} else {
				t = s
				pre = 0
			}
			if t > res {
				res = t
			}
			s = 0
		} else {
			validBlock = true
			if s+pre > res {
				res = s + pre
			}
		}
	}
	return res
}
