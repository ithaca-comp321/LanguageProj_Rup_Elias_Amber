package main

import (
	"fmt"
)

func insertLetter() {
	//string formatting

}

func drawBodyPart() {
	//list of stuff head, left arm, right arm, left leg, right leg, body,

}

func addGuessToSet() {
	//when letters is guessed then automatically add to set
	//if it already exists then return "enter another letter"

}

func chooseRandomWord(words) {
	////DOESN'T WORK
	randomWord = rand.string(len(words))

	pick := words[randomWord]
	return
}

func main() {

	//create a list of all the words that the player will guess

	//get the words inside of our defined list
	var words []string
	words = append(words, "bananas")
	words = append(words, "apples")
	words = append(words, "peach")
	words = append(words, "elephant")
	words = append(words, "orange")
	words = append(words, "burritos")

	//this is how we can print the whole list
	fmt.Printf("%v", words)

}

//randomly select word user has to guess and replace letters with '_'

//if they get the letter right, prompt to guess whole word, if wrong then add a body part

//create an empty set to hold all guessed letters that user entered
