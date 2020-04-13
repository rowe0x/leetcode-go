package leetcode

func maxDepthAfterSplit(seq string) []int {
	var depth int
	res := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			res[i] = depth % 2
			depth++
		} else {
			depth--
			res[i] = depth % 2
		}
	}
	return res
}