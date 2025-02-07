package main

import "fmt"

//go语言中切片的特性，切片为引用传递，在函数内部修改切片，上层切片的值也会改变
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		//向有序数组中插入key，从后向前找，比key大则后移一位
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

}
func main() {
	arr := []int{9, 5, 1, 4, 3}
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
