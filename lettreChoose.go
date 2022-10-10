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
	//fmt.Print(len(letterChoose))
	if len(letterChoose) > 2 {
		if wordToFind != letterChoose {
			return false
		} else {
			return true
		}
	} else {
		for _, valueWord := range wordToFind {
			for _, valueLettreChoose := range letterChoose {
				if string(valueWord) == string(valueLettreChoose) {
					return true
				}
			}
		}
	}

	return false
}

func VerifeChar(wordToFind string, wordUncomplet string) string {
	attempts := 11
	wordInProgresse := wordUncomplet
	for attempts > 2 {
		letterChoose := LettreChoose()
		if IsPresent(wordToFind, letterChoose) == true {
			if letterChoose == "e" {
				wordInProgresse = Reveal(wordToFind, wordInProgresse, "e")
				wordInProgresse = Reveal(wordToFind, wordInProgresse, "é")
				wordInProgresse = Reveal(wordToFind, wordInProgresse, "è")
			} else if letterChoose == "c" {
				wordInProgresse = Reveal(wordToFind, wordInProgresse, "c")
				wordInProgresse = Reveal(wordToFind, wordInProgresse, "ç")
			} else {
				wordInProgresse = Reveal(wordToFind, wordInProgresse, letterChoose)
			}
			fmt.Println(wordInProgresse)
			Position(attempts)
		} else {
			fmt.Println(Reveal(wordToFind, wordInProgresse, letterChoose))
			wordInProgresse = Reveal(wordToFind, wordInProgresse, letterChoose)
			attempts--
			fmt.Println(attempts)
			Position(attempts)
		}
		fmt.Print("\n")
		if wordInProgresse == wordToFind {
			return WinOrLoose(attempts, wordToFind)
		}
	}
	return WinOrLoose(attempts, wordToFind)
}

func Reveal(wordToFind string, wordInProgresse string, letterChoose string) string {
	word := []rune(wordInProgresse)
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
	fmt.Println("")
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

func WinOrLoose(attempts int, wordToFind string) string {
	var endPrint string
	if attempts == 0 {
		endPrint = "Dommage ! Vous avez perdu, le mot était :" + wordToFind
	} else {
		endPrint = "Bravo ! Vous avez gagné, le mot était :" + wordToFind
	}
	return endPrint
}
