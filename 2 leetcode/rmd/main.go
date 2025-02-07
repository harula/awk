package main

import "fmt"

func removeDup(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j++
		}
	}

	return j + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3}
	k := removeDup(nums)
	fmt.Println("k", k, "num:", nums)
}
