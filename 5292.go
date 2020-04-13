package leetcode

func isPossibleDivide(nums []int, k int) bool {
	if len(nums) % k != 0 {
		return false
	}
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	for len(freq) > 0 {
		var min = 1000000000
		for n := range freq {
			if n < min {
				min = n
			}
		}
		freq[min]--
		if freq[min] == 0 {
			delete(freq, min)
		}
		for i := 1; i < k; i++ {
			if freq[min+i] == 0 {
				return false
			}
			freq[min+i]--
			if freq[min+i] == 0 {
				delete(freq, min+i)
			}
		}
	}
	return true
}
