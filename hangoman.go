package main

import (
	"bufio"
	"fmt"
	"os"
	"pascha.com/resources"
	"strings"
)

var TPL = "  +---+\n  |   |\n  %s   |\n %s%s%s  |\n %s %s  |\n      |\n ===========\n"

func pickWord() (map[int32]struct{}, string, string) {
	country, city := resources.PickRandomCity()
	lettersToGuess := make(map[int32]struct{})
	for _, char := range strings.ToUpper(city) {
		lettersToGuess[char] = struct{}{}
	}
	if strings.ToUpper(city) == city {
		return lettersToGuess, strings.ToUpper(city), fmt.Sprintf("capital of %s", country)
	} else{
		return lettersToGuess, strings.ToUpper(city), fmt.Sprintf("city in %s", country)
	}
}

func printGallows(wrongGuessesCount int) {
	var gallows = []interface{}{"O", "/", "|", "\\", "/",  "\\"}
	for i:=5; i >= wrongGuessesCount; i-- {
		gallows[i] = " "
	}
	fmt.Printf(TPL, gallows...)
}

func charInSlice(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func printWord(wordToGuess string, guesses []int32) int {
	lettersLeft := 0
	var wordToPrint string
	for _, char := range wordToGuess {
		wordToPrint += " "
		if charInSlice(char, guesses) {
			wordToPrint += string(char)
		} else {
			wordToPrint += "_"
			lettersLeft += 1
		}
	}
	wordToPrint += "\n"
	fmt.Printf(wordToPrint)
	return lettersLeft
}

func printMisses(misses []int32, ) {
	var wordToPrint string
	for _, char := range misses {
		wordToPrint += " " + string(char)
	}
	wordToPrint += "\n"
	fmt.Printf(wordToPrint)
}

func printPrompt(){
	fmt.Printf("\nYour guess (just one letter): ")
}
func printCurrent(wordToGuess string, guesses []int32, misses []int32) int {
	printGallows(len(misses))
	lettersLeft := printWord(wordToGuess, guesses)
	printMisses(misses)
	return lettersLeft
}

func toUpper(char int32 ) int32 {
	return int32(strings.ToUpper(string(char))[0])
}
func readInput() int32 {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	return char
}

func main() {
	//var guesses []int32
	var guesses []int32
	var misses []int32
	lettersToGuess, wordToGuess, explanation := pickWord()
	fmt.Println("I picked a city... Can you guess what it is?")
	for {
		lettersLeft := printCurrent(wordToGuess, guesses, misses)
		if len(misses) == 6 {
			fmt.Printf("You lose... The city is %s (%s)\n", wordToGuess, explanation)
			return
		} else if lettersLeft == 0 {
			fmt.Printf("Congratulations! You guessed!\n")
			return
		}
		printPrompt()
		userInput := readInput()
		fmt.Print("\033[H\033[2J")
		if userInput == 27 {
			return
		}
		userInput = toUpper(userInput)
		if !charInSlice(userInput, guesses) {
			guesses = append(guesses, userInput)
			if _, ok := lettersToGuess[userInput]; !ok {
				misses = append(misses, userInput)
			}
		}

	}
}