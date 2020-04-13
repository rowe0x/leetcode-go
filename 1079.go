package leetcode

func numTilePossibilities(tiles string) int {
	freq := make([]int, 26)
	for i := 0; i < len(tiles); i++ {
		freq[tiles[i]-'A']++
	}
	var dfs func([]int) int
	dfs = func(freq []int) int {
		s := 0
		for i := 0; i < 26; i++ {
			if freq[i] == 0 {
				continue
			}
			s++
			freq[i]--
			s += dfs(freq)
			freq[i]++
		}
		return s
	}
	return dfs(freq)
}
