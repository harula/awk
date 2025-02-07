package main

import "fmt"

//input  [1,1,2,2,2,2,3,4,4,4]
//output 7 [1,1,2,2,3,4,4]
func rmvDup2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	count := 1
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++

		} else {
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 3, 4, 4, 4}
	k := rmvDup2(nums)
	fmt.Println("k", k, "nums:", nums)
}
