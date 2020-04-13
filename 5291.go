package leetcode

func findNumbers(nums []int) int {
	ans := 0
	for _, n := range nums {
		var count = 1
		n /= 10
		for n > 0 {
			count += 1
		}
		if count % 2 == 0 {
			ans++
		}
	}
	return ans
}