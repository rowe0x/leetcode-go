package leetcode

import (
	"math"
)

func minimumTotal(triangle [][]int) int {

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			up := len(triangle[i-1])
			val := math.MaxInt64
			if j-1 >= 0 {
				if val > triangle[i-1][j-1] {
					val = triangle[i-1][j-1]
				}
			}
			if j < up {
				if val > triangle[i-1][j] {
					val = triangle[i-1][j]
				}
			}
			triangle[i][j] += val
		}
	}

	var ans = triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if ans > triangle[len(triangle)-1][i] {
			ans = triangle[len(triangle)-1][i]
		}
	}
	return ans
}
