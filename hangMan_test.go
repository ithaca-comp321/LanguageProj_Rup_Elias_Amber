package main

import (
	"testing"
)

/*
//TODO test everything
func testTest(t *testing.T) {
	//example format
	result := 1 + 3

	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}*/

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestFormatWord(t *testing.T) {
	var word = "help"
	var j = "bsdvfejlfe"

	var newWord = formatWord(word)
	var newWord2 = formatWord(j)

	assertEqual(t, len(newWord), len(word)*2)

	assertEqual(t, len(newWord2), len(j)*2)

	assertEqual(t, newWord, "_ _ _ _ ")

}
