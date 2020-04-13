package leetcode

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var res int
	for i := 0; i <= len(calories)-k; i++ {
		var s int
		for j := i; j < i+k; j++ {
			s += calories[j]
		}
		if s < lower {
			res--
		}
		if s > upper {
			res++
		}
	}
	return res
}
