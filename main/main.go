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
	fmt.Println(firstOutcome)
	fmt.Println(hangman.FillHangman(word, firstOutcome))
}
