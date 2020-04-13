package leetcode

import "fmt"

func maxDistance(grid [][]int) int {
	var x1 int
	var y1 int
	landFound := false
	waterFound := false
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 && !landFound {
				x1, y1 = i, j
				landFound = true
			} else {
				waterFound = true
			}
			if landFound && waterFound {
				break
			}
		}
	}
	if !landFound || !waterFound {
		return -1
	}
	var visited = make([][]bool, len(grid))
	var validPos = func(x, y int) bool {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			return true
		}
		return false
	}
	for i := range grid {
		visited[i] = make([]bool, len(grid[0]))
	}
	var res int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !validPos(x, y) {
			return
		}
		if visited[x][y] {
			return
		}
		visited[x][y] = true
		if grid[x][y] == 0 {
			max := 1
			if validPos(x, y-1) && grid[x][y-1] > 1 {
				if max < grid[x][y-1]/2 {
					max = grid[x][y-1]/2
				}
				dfs(x, y-1)
			}
			if validPos(x, y+1) && grid[x][y+1] > 1 {
				if max < grid[x][y+1]/2 {
					max = grid[x][y+1]/2
				}
				dfs(x, y+1)
			}
			if validPos(x-1, y) && grid[x][y+1] > 1 {
				if max < grid[x-1][y]/2 {
					max = grid[x-1][y]/2
				}
				dfs(x-1, y)
			}
			if validPos(x+1, y) && grid[x+1][y] > 1 {
				if max < grid[x+1][y]/2 {
					max = grid[x+1][y]/2
				}
				dfs(x+1, y)
			}
			grid[x][y] =  max * 2
			if res < grid[x][y] / 2 {
				res = grid[x][y] / 2
			}
		}
	}
	dfs(x1, y1)
	return res
}


func maxDistance2(grid [][]int) int {
	landFound, waterFound := false, false
	var xl, yl int
	for i := 0; i < len(grid); i++ {
		if waterFound && landFound {
			break
		}
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && !landFound {
				landFound = true
				xl, yl = i, j
			} else if grid[i][j] == 0 && !waterFound {
				waterFound = true
			}
			if waterFound && landFound {
				break
			}
		}
	}
	if !waterFound || !landFound {
		return -1
	}
	ans := -1
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return
		}
		if grid[x][y] < 0 {
			return
		}
		if grid[x][y] != 1 {
			min := 0
			init := false
			for _, dx := range []int{-1, 1} {
				if x+dx >= 0 && x+dx < len(grid) {
					if grid[x+dx][y] == 1 {
						init = true
						break
					} else if grid[x+dx][y] < 0 {
						if grid[x+dx][y] < min {
							min = grid[x+dx][y]
						}
					}
				}
			}
			for _, dy := range []int{-1, 1} {
				if y+dy >= 0 && y+dy < len(grid[0]) {
					if grid[x][y+dy] == 1 {
						init = true
						break
					} else if grid[x][y+dy] < 0 {
						if grid[x][y+dy] < min {
							min = grid[x][y+dy]
						}
					}
				}
			}
			if init {
				grid[x][y] = -1
			} else {
				grid[x][y] += min-1
			}
			if ans < grid[x][y] {
				ans = grid[x][y]
			}
		}
		for _, dx := range []int{-1, 1} {
			x1 := x+dx
			dfs(x1, y)
		}
		for _, dy := range []int{-1, 1} {
			y1 := y+dy
			dfs(x, y1)
		}
	}
	fmt.Println(grid)
	dfs(xl, yl)
	fmt.Println(grid)
	return -ans
}
