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

func IsPresent(wordToFind string, lettreChoose string) bool { // func returne true if lettre choose by user is present in word to find
	if len(lettreChoose) > 1 && wordToFind == lettreChoose {
		return true
	}
	for _, valueWord := range wordToFind {
		for _, valueLettreChoose := range lettreChoose {
			if string(valueWord) == string(valueLettreChoose) {
				return true
			}
		}
	}
	return false
}

// func GameOver(attempts int) bool {
// 	if attempts == 0 {
// 		return true
// 	}
// 	return false
// }

func FillHangman(attempts int, wordToFind string, wordUncomplet string) {
	lettreChoose := LettreChoose()
	if attempts > 0 {
		if IsPresent(wordToFind, lettreChoose) == true {
			fmt.Println(Reveal(wordToFind, wordUncomplet, lettreChoose))
			wordUncomplet = Reveal(wordToFind, wordUncomplet, lettreChoose)
			fmt.Print("Is present")
		} else {
			fmt.Println(Reveal(wordToFind, wordUncomplet, lettreChoose))
			wordUncomplet = Reveal(wordToFind, wordUncomplet, lettreChoose)
			fmt.Print("Is not present")
			attempts--
		}
		fmt.Print("\n")
		fmt.Println(attempts)
		FillHangman(attempts, wordToFind, wordUncomplet)
	}
}

func Reveal(wordToFind string, wordUncomplet string, lettreChoose string) string {
	word := []rune(wordUncomplet)
	index := 0
	for _, letter := range wordToFind {
		for _, valueLettreChoose := range lettreChoose {
			if string(letter) == string(valueLettreChoose) {
				word[index] = rune(letter)
			}
		}
		index++
	}
	return string(word)
}
