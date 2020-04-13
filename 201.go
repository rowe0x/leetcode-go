package leetcode

import (
	"math"
)

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	lower := int(math.Log2(float64(m)))
	upper := int(math.Log2(float64(n)))
	if upper > lower {
		return 0
	}
	return 1 << uint(lower)
}
