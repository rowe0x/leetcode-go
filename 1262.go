package leetcode

import (
	"fmt"
	"sort"
)

func maxSumDivThree(nums []int) int {
	var reminder1 []int
	var reminder2 []int
	var ans int
	for _, n := range nums {
		switch n % 3 {
		case 0:
			ans += n
		case 1:
			reminder1 = append(reminder1, n)
		case 2:
			reminder2 = append(reminder2, n)
		}
	}
	fmt.Println(reminder1, reminder2)
	sort.Sort(sort.Reverse(sort.IntSlice(reminder1)))
	sort.Sort(sort.Reverse(sort.IntSlice(reminder2)))
	for len(reminder1) > 0 && len(reminder2) > 0 {
		if len(reminder1) < 3 {
			ans += reminder2[0] + reminder1[0]
			reminder2 = reminder2[1:]
			reminder1 = reminder1[1:]
		} else {
			t := reminder1[0] + reminder1[1] + reminder1[2]
			t1 := reminder1[0] + reminder2[0]
			if t > t1 {
				ans += t
				reminder1 = reminder1[3:]
			} else {
				ans += t1
				reminder1 = reminder1[1:]
				reminder2 = reminder2[1:]
			}
		}
	}
	for len(reminder1) >= 3 {
		ans += reminder1[0] + reminder1[1] + reminder1[2]
		reminder1 = reminder1[3:]
	}
	return ans
}
