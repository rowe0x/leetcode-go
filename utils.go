package leetcode

import (
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
