package leetcode

func minFallingPathSum(A [][]int) int {
	for r := 1; r < len(A); r++ {
		for c := 0; c < len(A[0]); c++ {
			t := A[r-1][c]
			if c-1 >= 0 {
				if t > A[r-1][c-1] {
					t = A[r-1][c-1]
				}
			}
			if c+1 < len(A[0]) {
				if t > A[r-1][c+1] {
					t = A[r-1][c+1]
				}
			}
			A[r][c] += t
		}
	}
	ans := A[len(A)-1][0]
	for _, v := range A[len(A)-1] {
		if ans > v {
			ans = v
		}
	}
	return ans
}
