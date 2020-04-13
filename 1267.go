package leetcode

func countServers(grid [][]int) int {
	row := make([]int, len(grid))
	col := make([]int, len(grid[0]))
	var ans int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				if row[r] == 1 {
					row[r] = 2 //  communicate found
				} else if row[r] == 0 {
					row[r] = 1
				}
				if col[c] == 1 {
					col[c] = 2
				} else if col[c] == 0 {
					col[c] = 1
				}
			}
		}
	}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				if row[r] == 2 || col[c] == 2 {
					ans++
				}
			}
		}
	}
	return ans
}