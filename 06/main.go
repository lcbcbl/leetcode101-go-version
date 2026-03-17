package main

// 70 爬楼梯
// 给定n节台阶，每次可以走一步或走两步，求一共有多少种方式可以走完这些台阶。

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
	}

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
