package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

}

func main() {
	arr := []int{5, 9, 6, 1, 4, 3, 10}
	selectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
