package leetcode

func setZeroes(matrix [][]int)  {
	return
	/*
	m := len(matrix)
	mm := 1 << m-1
	n := len(matrix[0])
	nm := 1 << n-1

	mbit, nbit := mm, nm
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				mbit &= mm >> (i+1) << (i+1) | ((1 << i)-1)
				nbit &= nm >> (i+1) << (i+1) | ((1 << i)-1)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mbit >> i | 0 == 0 {
				matrix[i][j] = 0
			}
			if nbit >> i | 0 == 0 {
				matrix[i][j] = 0
			}
		}
	}

	 */
}
