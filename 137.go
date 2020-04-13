package leetcode

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	pre := nums[0]
	preCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			preCount++
		} else if preCount == 1 {
				return pre
		} else {
			pre = nums[i]
			preCount++
		}
	}
	return pre
}
