package leetcode

func maxProfit(prices []int) int {
	var ans int
	count := len(prices)
	if count == 0 {
		return ans
	}
	var in int
	var isIn bool
	for i := 1; i < count; i++ {
		if prices[i] > prices[i-1] {
			if !isIn {
				isIn = true
				in = prices[i-1]
			}
		} else if isIn {
			ans += prices[i-1]-in
			isIn = false
		}
	}
	if isIn {
		ans += prices[count-1]-in
	}
	return ans
}