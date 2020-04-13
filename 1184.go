package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	s := distance[0]
	for i := 1; i < len(distance); i++ {
		s += distance[i]
		distance[i] += distance[i-1]
	}
	distance = append([]int{0},distance...)
	v := distance[destination] - distance[start]
	if v > s - v {
		v = s - v
	}
	return v
}
