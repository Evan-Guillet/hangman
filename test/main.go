package main

import (
	"fmt"
	"hangman"
)

func main() {
	if hangman.IsPresent("hello", hangman.LettreChoose()) == true {
		fmt.Print("is present")
	} else {
		fmt.Print("is not present")
	}

}
