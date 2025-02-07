package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 3, 0, 2, 2}
	result := majorityElement(nums)
	fmt.Println("The majority element is:", result)
}
