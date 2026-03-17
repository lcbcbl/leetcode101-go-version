package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	moreProfit := int64(0)
	for i := range k {
		if i >= k/2 {
			moreProfit += int64(prices[i])
		}
		moreProfit -= int64(prices[i] * strategy[i])
	}

	maxMoreProfit := int64(0)
	resI := -1
	for i := 0; i+k < n; i++ {
		if moreProfit > maxMoreProfit {
			maxMoreProfit = moreProfit
			resI = i
		}

		j := i + k + 1
		if j >= n {
			break
		}
		// 窗口左侧
		moreProfit -= int64(-prices[i] * strategy[i])
		// 窗口中间
		moreProfit -= int64(prices[i+1+k/2])
		// 窗口右侧
		moreProfit += int64(prices[j] - prices[j]*strategy[j])
	}

	ans := int64(0)
	for i := range prices {
		ans += int64(prices[i] * strategy[i])
	}
	if maxMoreProfit <= 0 {
		return ans
	}

	for i := range k {
		if i < k/2 {
			ans -= int64(prices[resI+i] * strategy[resI+i])
		} else {
			ans += int64(prices[resI+i] - prices[resI+i]*strategy[resI+i])
		}
	}

	return ans
}

func main() {
	// prices = [4,2,8], strategy = [-1,0,1], k = 2
	ans := maxProfit([]int{4, 2, 8}, []int{-1, 0, 1}, 2)

	fmt.Println(ans)
}
