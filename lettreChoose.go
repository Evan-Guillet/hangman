package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func LettreChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choisi une lettre :")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func IsPresent(word string, lettreChoose string) bool {
	if len(lettreChoose) > 1 && word == lettreChoose {
		return true
	}
	for _, valueWord := range word {
		for _, valueLettreChoose := range lettreChoose {
			if string(valueWord) == string(valueLettreChoose) {
				return true
			}
		}
	}
	return false
}
