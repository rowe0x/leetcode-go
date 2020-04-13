package leetcode

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]struct{}, 9)
	col := make([]map[byte]struct{}, 9)
	box := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		box[i] = make(map[byte]struct{})
		row[i] = make(map[byte]struct{})
		col[i] = make(map[byte]struct{})
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			v := board[i][j]
			if _, ok := row[i][v]; ok {
				return false
			}
			row[i][v] = struct{}{}

			if _, ok := col[j][v]; ok {
				return false
			}
			col[j][v] = struct{}{}

			r := i / 3
			boxIdx := r*3 + j/3
			if _, ok := box[boxIdx][board[i][j]]; ok {
				return false
			}
			box[boxIdx][board[i][j]] = struct{}{}
		}
	}
	return true
}
