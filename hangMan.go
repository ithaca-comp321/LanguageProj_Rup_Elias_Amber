package main

import (
	"fmt"
	"math/rand"
	"strings"

	//"regexp"
	"time"
)

//lazy global variables
var bodyParts = 0
var words = []string{"bananas", "apples", "peach", "elephant", "orange", "burritos", "ithaca", "college"}

//var currentBody []string
var guesses []string //create an empty set to hold all guessed letters that user entered
var wordToGuess string

//func drawBodyPart(wrongGuessCount int) []string {
//	//TODO broken, replaced with a counter for now
//	//for i := wrongGuessCount - 1; i <= wrongGuessCount; i++ {
//	//	currentBody = append(currentBody, bodyPart[i])
//	//}
//	//return currentBody
//
//	for i := 0; i <= wrongGuessCount; i++ {
//		currentBody[i] = bodyPart[i]
//	}
//	return currentBody
//}

func addGuessToSet(guess string) bool {
	//when letters is guessed then automatically add to set
	//if it already exists then return false, which means it was already guessed
	var duplicate = false
	for _, i := range guesses {
		if i == guess {
			duplicate = true
		}
	}
	if duplicate {
		return false
	} else {
		guesses = append(guesses, guess)
		return true
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

//current word starts off as the dashes
func getWordProgress(answer string, currentWord string, fullWord string) (string, bool) { //progress is not saved correctly
	//check if user letter is correct
	//replace "_" with the letter user guess right

	var new_word = ""
	var guessedRight = false

	//loop through dashes to place the letter user guessed
	for i := 0; i < len(currentWord); i++ {
		var letter = string(fullWord[i])
		if answer == letter {
			new_word += answer
			guessedRight = true
		} else {
			if string(currentWord[i]) != "_" {
				new_word += string(currentWord[i])
			} else {
				new_word += "_"
			}
		}
	}
	return new_word, guessedRight
}

func checkWholeWordGuess(guess string) bool {
	return strings.ToLower(guess) == strings.ToLower(wordToGuess)
}

func checkLose() bool {
	return bodyParts == 7
}

func main() { //user interface
	var wrongGuessCount int
	var answer string
	var currentWordProgress string
	var guessedRight bool

	fmt.Println("Welcome to Go Hangman.")
	var wordToGuess = chooseRandomWord() //randomly select word user has to guess
	currentWordProgress = formatWord(wordToGuess)
	fmt.Println("Your word is : " + currentWordProgress)
	fmt.Println("DEBUG: " + wordToGuess)
	currentWordProgress, _ = getWordProgress("0", currentWordProgress, wordToGuess) //the _ is a blank identifier, which ignores the second return value

	for answer != "quit" {
		fmt.Printf("Guesses so far: %v \n", guesses)
		fmt.Println("Enter 'letter' to guess a single letter or 'word' to guess the whole word: ")
		fmt.Scanf("%s", &answer)

		if answer == "letter" {
			fmt.Println("\nGuess a letter or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer)

			if len(answer) == 1 && addGuessToSet(strings.ToLower(answer)) { //TODO check answer is a character?
				currentWordProgress, guessedRight = getWordProgress(answer, currentWordProgress, wordToGuess) //get the results

				if !guessedRight { //if they didn't guess right
					wrongGuessCount++
					fmt.Println("That letter is not in the word.")
					//currentBody = drawBodyPart(wrongGuessCount)
					bodyParts++
					if checkLose() {
						fmt.Println("You lost!")
						break
					}
				}

				fmt.Println("Word Progress: " + currentWordProgress)
				if currentWordProgress == wordToGuess {
					fmt.Println("You win!")
					break
				}
				//fmt.Printf("Bodyparts visible: %v \n", currentBody)
				fmt.Println("Bodyparts visible: ", bodyParts)

			} else {
				fmt.Println("Please only enter a single letter, or make sure you didn't enter a duplicate guess.")
			}

		} else if answer == "word" {
			fmt.Println("\nGuess the word or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer) //need to worry about it not having only characters?
			//TODO
			if len(answer) == len(wordToGuess) {
				isCorrect := checkWholeWordGuess(strings.ToLower(answer))

				if isCorrect {
					fmt.Println("Congratulations! You got it right.")
					break
				} else {
					fmt.Println("Your guess was incorrect.")
					//currentBody = drawBodyPart(wrongGuessCount)
					bodyParts++
					//fmt.Printf("Bodyparts visible: %v \n", currentBody)
					fmt.Println("Bodyparts visible: ", bodyParts)
					if checkLose() {
						fmt.Println("You lost!")
						break
					}
				}

			} else {
				fmt.Println("The length of your guess doesn't match the length of the actual word, please try again.")
			}

		} else {
			fmt.Println("Please only enter 'letter' or 'word' exactly (or 'quit' to stop).")
		}
	}

	fmt.Println("Restart the program to play again")
}
