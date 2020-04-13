package  leetcode

func makeConnected(n int, connections [][]int) int {
	visited := make([]int, n)
	graph := make([][]int, 206)
	for _, connection := range connections {
		i := connection[0]
		j := connection[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if visited[i] == 1 {
			return 0
		}
		visited[i] = 1
		for n := range graph[i] {
			dfs(n)
		}
		return 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += dfs(i)
	}
	return ans - 1
}