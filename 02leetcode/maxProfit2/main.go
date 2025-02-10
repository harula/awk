package main

import (
	"fmt"
)

// 买入 卖出 7 1 2 3 4 5
func maxProfit2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	hold := -nums[0]
	notHold := 0
	//maxProfit := max(hold, notHold)
	for i := 1; i < len(nums); i++ {
		newHold := max(hold, notHold-nums[i])
		newNotHold := max(notHold, hold+nums[i])

		hold = newHold
		notHold = newNotHold
	}

	return notHold

}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 初始化两个变量，分别表示持有股票和不持有股票的最大利润
	hold := -prices[0]
	notHold := 0

	for i := 1; i < len(prices); i++ {
		// 更新不持有股票的最大利润
		newNotHold := max(notHold, hold+prices[i])
		// 更新持有股票的最大利润
		newHold := max(hold, notHold-prices[i])

		// 更新变量值
		hold = newHold
		notHold = newNotHold
	}

	return notHold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum profit:", maxProfit(prices)) // 输出最大利润
}
