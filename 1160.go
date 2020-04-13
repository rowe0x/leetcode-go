package leetcode

func countCharacters(words []string, chars string) int {
	var freq [26]int
	for i := 0; i < len(chars); i++ {
		freq[chars[i]-'a']++
	}
	var res int
	for _, word := range words {
		var reverse []uint8
		var good = true
		for i := 0; i < len(word); i++ {
			idx := word[i]-'a'
			if freq[idx] == 0 {
				good = false
				break
			}
			reverse = append(reverse, idx)
			freq[idx]--
		}
		if good {
			res += len(word)
		}
		for _, idx := range reverse {
			freq[idx]++
		}
	}
	return res
}
