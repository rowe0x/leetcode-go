package leetcode

func alphabetBoardPath(target string) string {
	var res []byte
	pos := [26][2]int{}
	for i := 0; i < 26; i++ {
		pos[i][0] = i / 5
		pos[i][1] = i % 5
	}
	x0, y0 := 0, 0
	for i := 0; i < len(target); i++ {
		t := target[i] - 'a'
		x, y := pos[t][0], pos[t][1]
		if x < x0 {
			for k := x; k < x0; k++ {
				res = append(res, 'U')
			}
		}
		if y < y0 {
			for k := y; k < y0; k++ {
				res = append(res, 'L')
			}
		}
		if x > x0 {
			for k := x0; k < x; k++ {
				res = append(res, 'D')
			}
		}
		if y > y0 {
			for k := y0; k < y; k++ {
				res = append(res, 'R')
			}
		}
		x0, y0 = x, y
		res = append(res, '!')
	}
	return string(res)
}
