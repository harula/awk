package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	for low < high {

		//从右往左找到第一个比povit小的值，移动到low的位置
		for low < high && arr[high] >= pivot {
			high--
		}

		arr[low] = arr[high]

		//从左往右找到第一个比pivot大的，移动到high的位置
		for low < high && arr[low] <= pivot {
			low++
		}

		arr[high] = arr[low]
	}

	//遍历结束后low右边的值都比pivot大，左边的值都比pivot小
	arr[low] = pivot

	fmt.Println("partition array:", arr, low)
	return low

}

func quickSort(arr []int, low int, high int) {

	if low < high {
		pivotPos := partition(arr, low, high)
		quickSort(arr, 0, pivotPos-1)
		quickSort(arr, pivotPos+1, high)
	}
}

func main() {
	arr := []int{5, 9, 6, 1, 4, 3, 10}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
