package main

import (
	"fmt"
	"math/rand"
	"strings"
	//"regexp"
	"time"
)

var bodyPart = []string{"head", "left arm", "right arm", "left leg", "right leg", "body"}
var words = []string{"bananas", "apples", "peach", "elephant", "orange", "burritos", "ithaca", "college"}

var currentBody []string
var guesses []string //create an empty set to hold all guessed letters that user entered
var word string

func drawBodyPart(wrongGuessCount int) []string {
	for i := wrongGuessCount; i <= len(bodyPart); i++ {
		currentBody = append(currentBody, bodyPart[i])
	}
	return currentBody
}

func addGuessToSet(guess string) string {
	//when letters is guessed then automatically add to set
	//if it already exists then return "enter another letter"
	var duplicate = false
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


func checkWholeWordGuess(guess string) bool {
	//checks if the user's guess for the whole word is exactly matching or not. make sure what is given is made lowercase

	return false
}

func checkLose() bool {
	return len(currentBody) == len(bodyPart)
}

func main() {

	var wrongGuessCount int
	var answer string

	fmt.Println("Welcome to Go Hangman.")
	word = chooseRandomWord() //randomly select word user has to guess
	dashes = formatWord(word)
	fmt.Println("Your word is : " + formatWord(word))
	getWordProgress(word)

	for answer != "quit" {

		fmt.Println("Enter 'letter' to guess a single letter or 'word' to guess the whole word: ")
		fmt.Scanf("%s", &answer)

		if answer == "letter" {
			fmt.Println("\nGuess a letter or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer)

			if len(answer) == 1 { //TODO check answer is a character?
				addGuessToSet(strings.ToLower(answer))
				wordProgress, guessedRight := getWordProgress() //get the results

				if !guessedRight { //if they didn't guess right
					wrongGuessCount++
					fmt.Println("That letter is not in the word.")
					currentBody = drawBodyPart(wrongGuessCount)
					if checkLose() {
						fmt.Println("You lost!")
						break
					}
				}

				fmt.Println("Word Progress: " + wordProgress)
				fmt.Printf("Bodyparts visible: %v \n", currentBody) //this prints the whole list?

			} else {
				fmt.Println("Please only enter a single letter")
			}

		} else if answer == "word" {
			fmt.Println("\nGuess the word or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer) //need to worry about it not having only characters?
			//TODO
			if len(answer) == len(word) {
				isCorrect := checkWholeWordGuess(strings.ToLower(answer))

				if isCorrect {
					fmt.Println("Congratulations! You got it right.")
					break
				} else {
					fmt.Println("Your guess was incorrect.")
					currentBody = drawBodyPart(wrongGuessCount)
					fmt.Printf("Bodyparts visible: %v \n", currentBody)
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
