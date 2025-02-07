package main

import "fmt"

//input [7,3,5,1,2,6,4]

func maxProfit(nums []int) int {
	maxProfit := 0
	minPrices := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < minPrices {
			minPrices = nums[i]
		} else if n := nums[i] - minPrices; n > maxProfit {
			maxProfit = n
		}
	}
	return maxProfit
}

func main() {
	nums := []int{7, 3, 5, 1, 2, 6, 4}
	m := maxProfit(nums)
	fmt.Println("m", m)
}
