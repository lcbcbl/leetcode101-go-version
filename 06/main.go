package main

// 70 爬楼梯
// 给定n节台阶，每次可以走一步或走两步，求一共有多少种方式可以走完这些台阶。
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 递归写法
// func climbStairs(n int) int {
// 	record := make([]int, n+1)

// 	var count func(n int) int
// 	count = func(n int) int {
// 		if n <= 2 {
// 			return n
// 		}

// 		if record[n] != 0 {
// 			return record[n]
// 		}

// 		ans := count(n-1) + count(n-2)
// 		record[n] = ans

// 		return ans
// 	}

// 	return count(n)
// }

// 198. 打家劫舍
func rob(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)
	dp[1] = nums[0]

	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}

	return dp[n]
}

// 413. 等差数列划分
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}

	ans := 0

	for i := range dp {
		ans += dp[i]
	}

	return ans
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}
