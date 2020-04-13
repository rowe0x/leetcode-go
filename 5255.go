package leetcode

func oddCells(n int, m int, indices [][]int) int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	for i := 0; i < len(indices); i++ {
		for c := 0; c < m; c++ {
			a[indices[i][0]][c]++
		}
		for r := 0; r < n; r++ {
			a[r][indices[i][1]]++
		}
	}
	var ans int
	for i := 0; i < len(indices); i++ {
		for j := 0; j < m; j++ {
			if indices[i][j] % 2 == 1 {
				ans++
			}
		}
	}
	return ans
}
