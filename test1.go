package main

import "testing"

func add(a, b int) int {
	return a + b
}

func testadd(t *testing.T) {

	r := add(2, 5)
	expected := 7
	if r != expected {
		t.Errorf("Expected %d but got %d", expected, r)
	}

}


