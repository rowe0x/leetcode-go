package leetcode

func numSmallerByFrequency(queries []string, words []string) []int {
	count := len(queries)
	res := make([]int, count)
	wordCount := len(words)
	wordFreq := make([]int, wordCount)
	for i, word := range words {
		wordFreq[i] = freqOfSmallChar(word)
	}
	var queryFreq int
	for i, query := range queries {
		var lessCount int
		queryFreq = freqOfSmallChar(query)
		for _, freq := range wordFreq {
			if queryFreq < freq {
				lessCount++
			}
		}
		res[i] = lessCount
	}
	return res
}

func freqOfSmallChar(s string) int {
	var freq [26]int
	var smallest uint8 = 26
	for i := 0; i < len(s); i++ {
		t := s[i] - 'a'
		if t <= smallest {
			smallest = t
			freq[smallest]++
		}
	}
	return freq[smallest]
}
