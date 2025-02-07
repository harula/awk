package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n             // 处理k值大于数组长度的情况
	reverse(nums, 0, n-1) // 反转整个数组
	reverse(nums, 0, k-1) // 反转前k个元素
	reverse(nums, k, n-1) // 反转剩余的元素
}

// reverse函数用于反转数组中从start到end的部分
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate2(nums []int, val int) {
	//nums2 := make([]int, val)
	//copy(nums2, nums[len(nums)-val:])
	//var nums2 []int
	//nums2 = append(nums2, nums[len(nums)-val:]...)

	val = val % len(nums)
	nums2 := make([]int, val)
	for i := 0; i < val; i++ {
		nums2[i] = nums[len(nums)-val+i]
	}

	for i := len(nums) - val - 1; i >= 0; i-- {
		nums[i+val] = nums[i]
	}
	for i := 0; i < val; i++ {
		nums[i] = nums2[i]
	}

}

//[7,6,5,4,3,2,1]
//[5,6,7,1,2,3,4]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 3
	rotate(nums, k)
	fmt.Println(nums) // 输出: [5 6 7 1 2 3 4]

	nums2 := []int{1}
	rotate2(nums2, k)
	fmt.Println(nums2)
}
