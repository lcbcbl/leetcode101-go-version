package main

// 01 背包问题
func package01() int {

	// n个物体，物体重量用weights数组表示 背包容量为c
	n := 3
	c := 5
	weights := []int{1, 2, 3}
	values := []int{1, 2, 3}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			w := weights[i-1]

			if w <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w]+values[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}

	}

	return dp[n][c]
}
