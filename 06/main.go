package main

import (
	"math"
)

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

// *****************
// 二维
// *****************
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

// 542. 01 矩阵
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// 左上 -> 右下
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}

			ans[i][j] = math.MaxInt - 1 // 避免 +1导致溢出
			if i > 0 {
				ans[i][j] = min(ans[i][j], ans[i-1][j]+1)
			}
			if j > 0 {
				ans[i][j] = min(ans[i][j], ans[i][j-1]+1)
			}
		}
	}
	// 右下 -> 左上
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}

			if i < m-1 {
				ans[i][j] = min(ans[i][j], ans[i+1][j]+1)
			}
			if j < n-1 {
				ans[i][j] = min(ans[i][j], ans[i][j+1]+1)
			}
		}
	}

	return ans
}

// 279. 完全平方数
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
//	1 <= n <= 1e4
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是
func numSquares(n int) int {
	nums := make([]int, 0, 100)
	for i := 1; i*i <= 1e4; i++ {
		nums = append(nums, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt

		for _, num := range nums {
			if i < num {
				break
			}

			dp[i] = min(dp[i], dp[i-num]+1)
		}
	}

	return dp[n]
}
