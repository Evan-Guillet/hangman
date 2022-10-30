package main

import (
	"hangman"
	"io/ioutil"
	
)

func main() {
	file, _ := ioutil.ReadFile("../words.txt")
	word := hangman.ChooseWord(file)
	firstOutcome := hangman.UncompletedWord(word)
	hangman.AsciiArt(firstOutcome)
	hangman.ResultDisplay(word, firstOutcome)
	hangman.AddWord(file)
}