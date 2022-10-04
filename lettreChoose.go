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
			Position(attempts)
		} else {
			fmt.Println(Reveal(wordToFind, wordUncomplet, lettreChoose))
			wordUncomplet = Reveal(wordToFind, wordUncomplet, lettreChoose)
			attempts--
			Position(attempts)
		}
		fmt.Print("\n")
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

func Position(i int) {
	fmt.Println("\n")
	switch i {
	case 11:
		fmt.Println("")
	case 10:
		fmt.Println("=========")
	case 9:
		fmt.Println(
			"       |  \n",
			"      |  \n",
			"      |  \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 8:
		fmt.Println(
			"  +---+  \n",
			"     |  \n",
			"     |  \n",
			"     |  \n",
			"     |  \n",
			"     |  \n",
			"=========",
		)
	case 7:
		fmt.Println(
			"   +---+  \n",
			"  |   | \n",
			"      |  \n",
			"      |  \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 6:
		fmt.Println(
			"   +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			"      |  \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 5:
		fmt.Println(
			"   +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			"  |   |  \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 4:
		fmt.Println(
			"   +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			" /|   | \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 3:
		fmt.Println(
			"   +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			" /|\\  |  \n",
			"      |  \n",
			"      |  \n",
			"=========",
		)
	case 2:
		fmt.Println(
			"  +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			" /|\\  |  \n",
			" /    |  \n",
			"      |  \n",
			"=========",
		)
	case 1:
		fmt.Println(
			"  +---+  \n",
			"  |   | \n",
			"  O   |  \n",
			" /|\\  |  \n",
			" / \\  |  \n",
			"      |  \n",
			"=========",
		)
	}

}