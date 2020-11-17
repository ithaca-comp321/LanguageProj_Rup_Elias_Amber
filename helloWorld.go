package main

import (
	"fmt"
)

func insertLetter() {
	//string formatting

}

func drawBodyPart(int count) {
	//list of stuff head, left arm, right arm, left leg, right leg, body,
    var bodyPart = []string {"head", "left arm", "right arm", "left leg", "right leg", "body"}
    
    string := RemoveIndex(helperstring []string, index int) []string {return append(s[:index], s[index+1:])}
    
    if count == 1{
        bobdyPart = RemoveIndex(bodyPart,1)
    }
    else{
        return bodyPart
    }
}

func addGuessToSet() {
	//when letters is guessed then automatically add to set
	//if it already exists then return "enter another letter"

}

func chooseRandomWord(words) {
	////DOESN'T WORK
	var randomWord = rand.string(len(words))

	pick := words[randomWord]
	return
}

func main() {

    //if they get the letter right, prompt to guess whole word, if wrong then add a body part

    fmt.Println("Enter a letter you will let to guess: ") 
  
    var guessLetter string 
    fmt.Scanf("%s",&guessLetter)

    if guessLetter == ...{
        fmt.Println("Do you wan a letter you will let to guess: ")
        var answer string 
        fmt.Scanf("%s",&answer) 
        if answer =="yes"{
            
        }
        
    }
    

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


//create an empty set to hold all guessed letters that user entered
