package leetcode

func numIslands(grid [][]byte) int {
	var ans int
	var dfs func(x, y int)
	var dirs = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dfs = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '#' || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '#'
		for _, dir := range dirs {
			dfs(x+dir[0], y+dir[1])
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}

		}
	}
	return ans
}
