package leetcode

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	var jc, sc int
	t := tomatoSlices - 2 * cheeseSlices
	if t < 0 {
		return []int{}
	}
	if t % 2 == 1 {
		return []int{}
	}
	jc = t/2
	sc = cheeseSlices-jc
	return []int{jc, sc}
}
