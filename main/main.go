package main

import (
	"io/ioutil"
	"hangman"
	"fmt"
)

func main () {
	file,_ := ioutil.ReadFile("../words.txt")
	word := hangman.ChooseWord(file)
	firstOutcome := hangman.UncompletedWord(word)
	fmt.Println(firstOutcome)
	attempts := 10
	hangman.FillHangman(attempts, word, firstOutcome)

}