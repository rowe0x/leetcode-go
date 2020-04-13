package leetcode

import (
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min := arr[len(arr)-1]-arr[0]
	candidate := make(map[int][][]int)
	var diff int
	for i := 1; i < len(arr); i++ {
		diff = arr[i]-arr[i-1]
		if diff <= min {
			min = diff
			candidate[diff] = append(candidate[diff], []int{arr[i-1], arr[i]})
		}
	}
	return candidate[min]
}