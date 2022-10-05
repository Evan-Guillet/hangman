package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func LettreChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choisi une letter :")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func IsPresent(wordToFind string, letterChoose string) bool { // func returne true if letter choose by user is present in word to find
	if len(letterChoose) > 1 && wordToFind == letterChoose {
		return true
	}
	for _, valueWord := range wordToFind {
		for _, valueLettreChoose := range letterChoose {
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

func FillHangman( wordToFind string, wordUncomplet string) string {
	attempts := 11
	for attempts > 0 {
		letterChoose := LettreChoose()
		if IsPresent(wordToFind, letterChoose) == true {
			fmt.Println(Reveal(wordToFind, wordUncomplet, letterChoose))
			wordUncomplet = Reveal(wordToFind, wordUncomplet, letterChoose)
			Position(attempts)
		} else {
			fmt.Println(Reveal(wordToFind, wordUncomplet, letterChoose))
			wordUncomplet = Reveal(wordToFind, wordUncomplet, letterChoose)
			attempts--
			fmt.Println(attempts)
			Position(attempts)
		}
		fmt.Print("\n")
		if wordUncomplet == wordToFind {
			return WinOrLoose(attempts, wordToFind)
		}
	}
	return WinOrLoose(attempts,wordToFind)
}

func Reveal(wordToFind string, wordUncomplet string, letterChoose string) string {
	word := []rune(wordUncomplet)
	index := 0
	for _, letter := range wordToFind {
		for _, valueLettreChoose := range letterChoose {
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

func WinOrLoose (attempts int, wordToFind string ) string {
	var endPrint string
	if attempts == 0 {
		endPrint = "Dommage ! Vous avez perdu, le mot était :" + wordToFind
	} else {
		endPrint = ("Bravo ! Vous avez gagné, le mot était :" + wordToFind)
	}
	return endPrint
}