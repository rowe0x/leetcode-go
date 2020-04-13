package leetcode

func minimumSwap(s1 string, s2 string) int {
	var ans int
	var x1, x2, y1, y2 int
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			x1++
			y2++
			if x1 == 2 && y2 == 2 {
				x1 -= 2
				y2 -= 2
				ans++
			}
		} else {
			y1++
			x2++
			if y1 == 2 && x2 == 2 {
				y1 -= 2
				x2 -= 2
				ans++
			}
		}
	}
	if (x1+x2) % 2 == 1 || (y1+y2) % 2 == 1 {
		return -1
	}
	if x1 == 1 || y1 == 1 {
		return -1
	}
	ans++
	return ans
}
