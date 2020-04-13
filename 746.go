package leetcode

func minCostClimbingStairs(cost []int) int {
	count := len(cost)
	var a, b = cost[0], cost[1]
	var min int
	for i := 2; i < count; i++ {
		min = a
		if a > b {
			min = b
		}
		a = b
		b = min + cost[i]

	}
	if a > b {
		return b
	}
	return a
}

