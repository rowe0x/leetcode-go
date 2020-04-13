package leetcode

func isBoomerang(points [][]int) bool {
	var dx1,dx2,dy1,dy2 int
	dx1 = points[1][0] - points[0][0]
	dy1 = points[1][1] - points[0][1]

	dx2 = points[2][0] - points[1][0]
	dy2 = points[2][1] - points[1][1]

	if dx1 * dy2 == dy1 * dx2 {
		return false
	}
	return true
}