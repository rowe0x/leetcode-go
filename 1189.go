package leetcode

func maxNumberOfBalloons(text string) int {
	var freq [26]int
	for i := 0; i < len(text); i++ {
		freq[text[i]-'a']++
	}
	var ans int
	bFreq := freq['b'-'a']
	aFreq := freq['a'-'a']
	lFreq := freq['l'-'a']
	oFreq := freq['o'-'a']
	nFreq := freq['n'-'a']
	if ans > bFreq {
		ans = bFreq
	}
	if ans > aFreq {
		ans = aFreq
	}
	if ans > lFreq/2 {
		ans = lFreq/2
	}
	if ans > oFreq/2 {
		ans = oFreq/2
	}
	if ans > nFreq {
		ans = nFreq
	}
	return ans
}
