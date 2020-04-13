package leetcode

func sortArray(nums []int) []int {
	reg := make([]int, len(nums))
	mergeSort(nums, reg, 0, len(nums)-1)
	return nums
}

func mergeSort(nums, reg []int, start, end int) {
	if start >= end {
		return
	}
	m := start + (end-start)/2
	start1, end1 := start, m
	start2, end2 := m+1, end
	mergeSort(nums, reg, start1, end1)
	mergeSort(nums, reg, start2, end2)
	k := start
	for start1 <= end1 && start2 <= end2 {
		if nums[start1] < nums[start2] {
			reg[k] = nums[start1]
			start1++
		} else {
			reg[k] = nums[start2]
			start2++
		}
		k++
	}
	for start1 <= end1 {
		reg[k] = nums[start1]
		k++
		start1++
	}
	for start2 <= end2 {
		reg[k] = nums[start2]
		k++
		start2++
	}
	for i := start; i <= end; i++ {
		nums[i] = reg[i]
	}
}
