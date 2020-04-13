package leetcode

func tictactoe(moves [][]int) string {
	board := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]byte, 3)
	}
	row := make([]int, 3)
	col := make([]int, 3)
	diagonal := make([]int, 2)
	inATurn := true
	for r := 0; r < len(moves); r++ {
		if inATurn {
			board[moves[r][0]][moves[r][1]] = 'X'
			row[moves[r][0]] += 1
			if row[moves[r][0]] == 3 {
				return "A"
			}
			col[moves[r][1]] += 1
			if col[moves[r][1]] == 3 {
				return "A"
			}
			if moves[r][0] == moves[r][1] {
				diagonal[0] += 1
				if diagonal[0] == 3 {
					return "A"
				}
			}
			if moves[r][0] + moves[r][1] == 2 {
				diagonal[1]	+= 1
				if diagonal[1] == 3 {
					return "A"
				}
			}
			inATurn = false
		} else {
			board[moves[r][0]][moves[r][1]] = 'O'
			row[moves[r][0]] -= 1
			if row[moves[r][0]] == -3 {
				return "B"
			}
			col[moves[r][1]] -= 1
			if col[moves[r][1]] == -3 {
				return "B"
			}
			if moves[r][0] == moves[r][1] {
				diagonal[0] -= 1
				if diagonal[0] == -3 {
					return "B"
				}
			}
			if moves[r][0] + moves[r][1] == 2 {
				diagonal[1]	-= 1
				if diagonal[1] == -3 {
					return "B"
				}
			}
			inATurn = true
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
