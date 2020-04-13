package leetcode

func countSquares(matrix [][]int) int {
	rc, cc := len(matrix), len(matrix[0])
	var ans int
	for r := 1; r < rc; r++ {
		for c := 1; c < cc; c++ {
			if matrix[r][c] == 0 {
				continue
			}
			lv := matrix[r][c-1]
			uv := matrix[r-1][c]
			luv := matrix[r-1][c-1]
			var min = lv
			if uv < min {
				min = uv
			}
			if luv < min {
				min = luv
			}
			if min > 0 {
				matrix[r][c] = min+1
			}
		}
	}
	for r := 0; r < rc; r++ {
		for c := 0; c < cc; c++ {
			ans += matrix[r][c]
		}
	}
	return ans
}
