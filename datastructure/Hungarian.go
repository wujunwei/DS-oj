package datastructure

func maximumInvitations(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	pre := make([]int, m+n)
	for i := 0; i < len(pre); i++ {
		pre[i] = -1
	}
	edges := make([][]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				edges[i] = append(edges[i], j+m)
				edges[j+m] = append(edges[j+m], i)
			}
		}
	}
	var vis map[int]bool
	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, j := range edges[i] {
			if vis[j] {
				continue
			}
			vis[j] = true
			if pre[j] == -1 || dfs(pre[j]) {
				pre[j] = i
				return true
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		vis = make(map[int]bool)
		if dfs(i) {
			ans++
		}
	}
	return ans
}
