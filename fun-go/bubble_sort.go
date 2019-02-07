package main

import (
	"math/rand"
	"fmt"
)

func generateArray() []int {
	var arr []int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}
	return arr
}

func bubbleSort(arr []int) []int {
	var swapped bool
	for i := 0; i < len(arr) - 1; i++ {
		swapped = false
		for j := 0; j < len(arr) - i - 1; j++ {
			if (arr[j] > arr[j + 1]) {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp 
				swapped = true
			}
		}
		if swapped == false {
			break 
		}
	}
	return arr 
}

func main() {
	arr := []int{7, 4, 3, 6, 5, 1, 2, 9, 8}
	arr = bubbleSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

}
