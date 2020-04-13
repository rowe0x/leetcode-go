package leetcode

func minCostToMoveChips(chips []int) int {
	var divisibleCount int
	var nonDivisibleCount int
	for _, c := range chips {
		if c % 2 == 0 {
			divisibleCount++
		} else {
			nonDivisibleCount++
		}
	}
	if divisibleCount > nonDivisibleCount {
		return nonDivisibleCount
	}
	return divisibleCount
}
