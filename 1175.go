package leetcode

func numPrimeArrangements(n int) int {
	if n <= 1 {
		return 0
	}
	var res = 1
	var primeCount int
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}
	for i := primeCount; i > 0; i-- {
		res = (res * i) % (1000000000 + 7)
	}
	for i := n-primeCount; i > 0; i-- {
		res = (res * i) % (1000000000 + 7)
	}
	return res
}
