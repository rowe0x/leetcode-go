package leetcode

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var s int
	var twoIdx []int
	for i, n := range colsum {
		s += n
		if n == 2 {
			twoIdx = append(twoIdx, i)
		}
	}
	if s != (upper+lower) {
		return nil
	}
	res := make([][]int, 2)
	for i := 0; i < 2; i++ {
		res[i] = make([]int, len(colsum))
	}
	for i := 0; i < len(twoIdx); i++ {
		res[0][twoIdx[i]] = 1
		upper--
		res[1][twoIdx[i]] = 1
		lower--
		colsum[twoIdx[i]] = 0
	}
	for _, c := range colsum {
		if c == 0 {
			continue
		}
		if upper > 0 {
			res[0][c] = 1
		} else {
			res[1][c] = 1
		}
	}
	return res
}
