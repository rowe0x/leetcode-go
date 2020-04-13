package leetcode

func canMakePaliQueries(s string, queries [][]int) []bool {
	var res []bool
	canMakePaliQuery := func(query []int) bool {
		left, right, up := query[0], query[1], query[2]
		freq := make(map[byte]int)
		for i := left; i <= right; i++ {
			freq[s[i]]++
		}
		var oddCount int
		for _, v := range freq {
			if v % 2 == 1 {
				oddCount++
			}
		}
		if oddCount % 2 == 1 {
			oddCount--
		}
		return up*2 >= oddCount
	}
	for _, query := range queries {
		res = append(res, canMakePaliQuery(query))
	}
	return res
}

