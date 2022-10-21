package main

import (
	"fmt"
	"hangman"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("../words.txt")
	word := hangman.ChooseWord(file)
	firstOutcome := hangman.UncompletedWord(word)
	hangman.AsciiArt(firstOutcome)
	fmt.Println(hangman.VerifeChar(word, firstOutcome))
}
