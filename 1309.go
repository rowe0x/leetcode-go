package leetcode

func freqAlphabets(s string) string {
	var ans []byte
	var m = make(map[int]struct{})
	for i := 0;i < len(s); i++ {
		if s[i] == '#' {
			m[i-2] = struct{}{}
		}
	}
	for i := 0; i < len(s); {
		if s[i] == '#' {
			continue
		}
		if _, ok := m[i]; ok {
			ans = append(ans, 10 * (s[i]-'0') + (s[i+1]-'0') - 1 + 'a')
			i++
		} else {
			ans = append(ans, s[i]-'1'+'a')
		}
		i++
	}
	return string(ans)
}
