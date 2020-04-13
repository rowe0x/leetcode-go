package leetcode

func defangIPaddr(address string) string {
	count := len(address)
	res := make([]byte, count)
	for i := 0; i < count; i++ {
		if address[i] != '.' {
			res = append(res, address[i])
		} else {
			res = append(res, '[', '.', ']')
		}
	}
	return string(res)
}
