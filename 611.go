package leetcode

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var ans int
	var combination func(cur []int, start int)
	combination = func(cur []int, start int) {
		if len(cur) == 3 {
			if isTriangle(cur) {
				ans++
			}
			return
		}
		for i := start; i <= len(nums)-3; i++ {
			cur = append(cur, nums[i])
			combination(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	combination(nil, 0)
	return ans
}

func isTriangle(nums []int) bool {
	sort.Ints(nums)
	if nums[0]+nums[1] > nums[2] && nums[2]-nums[1] < nums[0] {
		return true
	}
	return false
}
