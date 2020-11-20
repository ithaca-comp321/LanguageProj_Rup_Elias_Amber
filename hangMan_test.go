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

func TestGetWordProgress(t *testing.T) {
	var word = "banana"
	var hide = formatWord(word)

	var guessedWord, _ = getWordProgress("b", hide, word)

	assertEqual(t, guessedWord, " b _  _  _  _  _  ")

	var guessedWord1, _ = getWordProgress("a", guessedWord, word)
	assertEqual(t, guessedWord1, "b a _ _ _ _ ")

	var guessedWord2, _ = getWordProgress("k", guessedWord1, word)
	assertEqual(t, guessedWord2, "b a _ _ _ _ ")

}

func TestAddGuessToSet(t *testing.T) {
	
	//ture or false 

	var letter = "a"
	var placeLetter = addGuessToSet(letter)

	assertEqual(t, placeLetter, true)

}
