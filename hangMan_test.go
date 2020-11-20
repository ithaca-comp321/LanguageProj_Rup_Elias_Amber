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

	assertEqual(t, len(newWord), len(word))

	assertEqual(t, len(newWord2), len(j))

	assertEqual(t, newWord, "____")

}

func TestGetWordProgress(t *testing.T) {
	var word = "banana"
	var hide = formatWord(word)

	var guessedWord, _ = getWordProgress("b", hide, word)

	assertEqual(t, guessedWord, "b_____")

	var guessedWord1, _ = getWordProgress("a", guessedWord, word)
	assertEqual(t, guessedWord1, "ba_a_a")

	var guessedWord2, _ = getWordProgress("k", guessedWord1, word)
	assertEqual(t, guessedWord2, "ba_a_a")

}

func TestAddGuessToSet(t *testing.T) {

	//ture or false

	var letter = "a"
	var placeLetter = addGuessToSet(letter)
	assertEqual(t, placeLetter, true)

	placeLetter = addGuessToSet(letter)
	assertEqual(t, placeLetter, false)
}
