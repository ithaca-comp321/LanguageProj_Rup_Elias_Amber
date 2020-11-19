package main

import (
	"fmt"
	"math/rand"

	//"regexp"
	"time"
)

var bodyPart = []string{"head", "left arm", "right arm", "left leg", "right leg", "body"}
var words = []string{"bananas", "apples", "peach", "elephant", "orange", "burritos", "ithaca", "college"}

var currentBody []string
var guesses []string //create an empty set to hold all guessed letters that user entered
var word string

func drawBodyPart(wrongGuessCount int) []string {
	for i := 0; i <= wrongGuessCount; i++ {
		currentBody = append(currentBody, bodyPart[i])
	}
	return currentBody
}

func addGuessToSet(guess string) string {
	//when letters is guessed then automatically add to set
	//if it already exists then return "enter another letter"
	var duplicate bool = false
	for _, i := range guesses {
		if i == guess {
			duplicate = true
		}
	}
	if duplicate {
		return "You guessed a duplicate letter, guess another.\n"
	} else {
		guesses = append(guesses, guess)
		return "You guessed " + guess + "\n"
	}
}

func chooseRandomWord() string {
	rand.Seed(time.Now().Unix())

	//var randomIndx = rand.Intn(len(words))

	var pick = words[rand.Intn(len(words))]

	return pick
}

func formatWord(word string) string {
	var dashes = ""

	for i := 0; i < len(word); i++ {
		dashes += "_"

	}
	return dashes
}

func getWordProgress() (answer string, dashes string, fullWord string) { //TODO
	//check if user letter is correct 
	//replace "_" with the letter user guess right

	var newdashes = "";

	//loop through dashs to place the letter user guess 
	for i, r := range dashes {
			var letter = string(word[i])
			if answer == letter {
				newdashes += answer
			} else {
				newdashes += "_"
			}
		}
		return answer, fullWord, newdashes
	}


func main() {

	var wrongGuessCount int
	var answer string

	word = chooseRandomWord() //randomly select word user has to guess
	dashes = formatWord(word)
	fmt.Println("Your word is : " + formatWord(word))
	getWordProgress(word)

	for answer != "quit" {
		//TODO Ask if they want to guess a letter or the whole word
		fmt.Scanf("%s", &answer)

		if answer == "letter" {
			fmt.Println("\nGuess a letter or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer) //TODO need to do some user error checking here, make sure they only enter a letter

			addGuessToSet(answer)
			wordProgress, guessedRight := getWordProgress()

			if !guessedRight { //if they didn't guess right
				wrongGuessCount++
				currentBody = drawBodyPart(wrongGuessCount)
			}

			fmt.Println("Word Progress: " + wordProgress)
			fmt.Printf("Bodyparts visible: %v", currentBody) //this prints the whole list?
		} else if answer == "word" {
			//TODO, also if their guessed word length doesn't match the given word length, make em guess again
		} else {
			//TODO user is dumb
		}
	}
}
