package leetcode

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x1,y1,x2,y2 := rec1[0],rec1[1],rec1[2],rec1[3]
	a1,b1,a2,b2 := rec2[0],rec2[1],rec2[2],rec2[3]
	// contain
	if x1 <= a1 && y1 <= b1 && x2 >= a2 && y2 >= b2 {
		return true
	}
	if x1 >= a1 && y1 >= b1 && x2 <= a2 && y2 <= b2 {
		return true
	}
	// bottom left
	if a1 > x1 && x2 > a1 && b1 > y1 && y2 > b1 {
		return true
	}
	// up right
	if a2 > x1 && a2 < x2 && b2 > y1 && b2 < y2 {
		return true
	}
	// up left
	if x2 > a1 && a1 < x1 && y1 > b1 && y1 < b2 {
		return true
	}
	// bottom right
	if a2 > x1 && a2 < x2 && b1 > y1 && b1 < y2 {
		return true
	}
	return false
}
