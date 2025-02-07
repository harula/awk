package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr) - 1
	for i := 0; i < n; i++ {

		swapped := false
		//后i个元素已经排过序
		for j := 0; j < n-i; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}

		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{9, 5, 1, 4, 3}
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
