package main

import (
	"testing"
)

//TODO test everything
func testTest(t *testing.T) {
	//example format
	result := 1 + 3

	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}
