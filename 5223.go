package leetcode

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var ans [][]int
	var cb [8][8]int
	for _, queen := range queens {
		cb[queen[0]][queen[1]] = 'Q'
	}
	for l := king[1]-1; l >= 0; l-- {
		if cb[king[0]][l] == 'Q' {
			ans = append(ans, []int{king[0], l})
			break
		}
	}
	for r := king[1]+1; r < 8; r++ {
		if cb[king[0]][r] == 'Q' {
			ans = append(ans, []int{king[0], r})
			break
		}
	}
	for u := king[0]-1; u >= 0; u-- {
		if cb[king[1]][u] == 'Q' {
			ans = append(ans, []int{u, king[1]})
			break
		}
	}
	for d := king[0]+1; d < 8; d++ {
		if cb[king[1]][d] == 'Q' {
			ans = append(ans, []int{d, king[1]})
			break
		}
	}
	for x, y := king[0]-1, king[1]-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if cb[x][y] == 'Q' {
			ans = append(ans, []int{x, y})
			break
		}
	}
	for x, y := king[0]+1, king[1]-1; x < 8 && y >= 0; x, y = x+1, y-1 {
		if cb[x][y] == 'Q' {
			ans = append(ans, []int{x, y})
			break
		}
	}
	for x, y := king[0]-1, king[1]+1; x >=0 && y < 8; x, y = x-1, y+1 {
		if cb[x][y] == 'Q' {
			ans = append(ans, []int{x, y})
			break
		}
	}
	for x, y := king[0]+1, king[1]+1; x < 8 && y < 8; x, y = x+1, y+1 {
		if cb[x][y] == 'Q' {
			ans = append(ans, []int{x, y})
			break
		}
	}
	return ans
}
