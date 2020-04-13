package leetcode

func numTimesAllBlue(light []int) int {
	var ans int
	count,  max := 0,  1
	for _, idx := range light {
		if idx > max {
			max = idx
		}
		count++
		if max == count {
			ans++
		}
	}
	return ans
}
