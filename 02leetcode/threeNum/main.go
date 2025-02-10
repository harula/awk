package main

import (
	"fmt"
	"sort"
)

func threeNum(nums []int) [][]int {
	sort.Ints(nums)
	rst := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				rst = append(rst, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				// -3 -2 -1 0 1 2
				left++
			} else if sum > 0 {
				// -3 0 1 2 3 4
				right--
			}
		}
	}

	return rst
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	rst := threeNum(nums)
	fmt.Println("rst :", rst)
}
