package main

// 695. 岛屿的最大面积 https://leetcode.cn/problems/max-area-of-island/description/
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}

		if grid[i][j] != 1 {
			return 0
		}

		// 2 标记已搜索
		grid[i][j] = 2
		area := 1
		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}
		for next := range dx {
			area += dfs(i+dx[next], j+dy[next])
		}

		return area
	}

	ans := 0
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}

	return ans
}

// 547. 省份数量 https://leetcode.cn/problems/number-of-provinces/description/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	traveled := make([]bool, n)

	var dfs func(city int)
	dfs = func(city int) {
		traveled[city] = true

		for next, isCon := range isConnected[city] {
			if isCon == 1 && !traveled[next] {
				dfs(next)
			}
		}
	}

	ans := 0
	for i := range n {
		if traveled[i] {
			continue
		}

		dfs(i)
		ans++
	}

	return ans
}
