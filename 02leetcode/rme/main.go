package main

import "fmt"

func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{1, 2, 3, 2, 1}
	k := removeElement(nums, 2)
	fmt.Println("k:", k, "nums:", nums[:k])
}
