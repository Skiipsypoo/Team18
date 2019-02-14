package main

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int {9, 8, 7, 6, 5, 4}

	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] < arr[i + 1] {
			t.Errorf("Incorrect sorting")
		}
	}
}

