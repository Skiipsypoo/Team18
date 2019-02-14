package main

import "testing"

func sumTest(t *testing.T) {
	if sum(5, 5) != 10 {
		t.Errorf("Inc")
	}
}