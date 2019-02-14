package main

import "fmt"

func quickSort(arr[] int) [] int{
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

}