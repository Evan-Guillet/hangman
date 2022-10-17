package hangman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const colorRed = "\033[1;31m"
const colorGreen = "\033[1;32m"
const colorReset = "\033[0m"

func LetterChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choose a letter :")
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')

	if error == nil {
		if text == "é\n" || text == "è\n" {
			text = "e\n"
		}
		if text == "ç\n" {
			text = "c\n"
		}
	}
	return text
}
func IsPresent(wordToFind string, letterChoose string) bool { // func returne true if letter choose by user is present in word to find
	if len(letterChoose) > 1 {
		for ii, valueLetterChoose := range letterChoose {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valueLetterChoose {
					return false
				}
			}
		}
		return true
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

func AlreadySaid(letterChoose string, wordSaid string) string {
	err := false
	var said []string
	if len(letterChoose) == 1 {
		for _, valueWordSaid := range wordSaid {
			if letterChoose == string(valueWordSaid) {
				return wordSaid
			}
		}
	} else {
		for _, valueLetterChoose := range letterChoose {
			if valueLetterChoose < 'a' || valueLetterChoose > 'z' {
				err = true
				
			}
		}
		if !err {
			wordSaid = wordSaid + letterChoose
			said = strings.Split(wordSaid, "\n")
			wordString := strings.Join(said, " ")
			fmt.Println("Already tried :", wordString)
		} else {
			fmt.Println(string(colorRed), "You entered an invalid letter", string(colorReset))
		}
	}
	return wordSaid
}

func VerifeChar(wordToFind string, wordUncomplet string) string {
	isSaid := false
	attempts := 11
	var wordSaid string
	wordInProgresse := wordUncomplet
	for attempts > 1 {
		letterChoose := LetterChoose()
		fmt.Println()
		said := []rune(wordSaid)
		letter := []rune(letterChoose)
		
		for i := 0; i < len(wordSaid); i++ {
			isSaid = false
			if !isSaid {
				if letter[0] == said[i] {
					isSaid = true
					break
				}
				
			}
			
		}
		
		wordSaid = AlreadySaid(letterChoose, wordSaid)
		letterChoose = strings.Replace(letterChoose, "\n", "", -1)
		wordInProgresse = Reveal(wordToFind, wordInProgresse, letterChoose)
		AsciiArt(wordInProgresse)
		fmt.Println()
		

		if !isSaid {	
			if IsPresent(wordToFind, letterChoose) {
				Position(attempts)
				fmt.Println(string(colorGreen), "__________________________________________", string(colorReset))
				
			} else {
				attempts--
				Position(attempts)
				fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
			}
		} else {
			fmt.Println(string(colorRed), "You already choose this letter", string(colorReset))
			attempts--
			Position(attempts)
			fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
		}
		
		fmt.Println()
		fmt.Println("remaining try :", attempts-1)
		fmt.Println("\n\n")
		if wordInProgresse == wordToFind {
			return WinOrLoose(attempts, wordToFind)
		}
	}
	return WinOrLoose(attempts, wordToFind)
}

func Reveal(wordToFind string, wordInProgresse string, letterChoose string) string {
	word := []rune(wordInProgresse)
	index := 0
	if len(letterChoose) > 1 {
		for ii, valueLetterChoose := range letterChoose {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valueLetterChoose {
					return wordInProgresse
				}
			}
		}
		return wordToFind
	} else {
		for _, letter := range wordToFind {
			for _, valueLettreChoose := range letterChoose {
				if string(letter) == string(valueLettreChoose) {
					word[index] = rune(letter)
				}
			}
			index++
		}
	}
	return string(word)
}

func Position(attempts int) {
	hangFile, _ := ioutil.ReadFile("../position_hangman.txt")
	file := strings.Split(string(hangFile), ",,")
	position := 11 - attempts
	if position < 0 {
		fmt.Println(file[0])
	} else {
		fmt.Println(file[position])
	}

}

func WinOrLoose(attempts int, wordToFind string) string {
	var endPrint string
	if attempts == 1 {
		endPrint = "Dommage ! Vous avez perdu, le mot était :" + wordToFind
	} else {
		endPrint = "Bravo ! Vous avez gagné, le mot était :" + wordToFind
	}
	return endPrint
}


func AsciiArt(wordUncomplet string) {
	fileIncome,_ := ioutil.ReadFile("../standard.txt")
	file := strings.Split(string(fileIncome), ",,")
	var letter []string
	word := []rune(wordUncomplet)
	var art string
	for j := 0; j < 8; j++ {
		for k := 0; k < len(word); k++ {
			if word[k] >= 'a' && word[k] <= 'z'{
				difference := int(word[k]) - int('a')
				letter = strings.Split(string(file[33+difference]), "\n")
				art += strings.Replace(letter[j], "\r", "", -1)
			} else {
				letter = strings.Split(string(file[63]), "\n")
				art += strings.Replace(letter[j], "\r", "", -1)
			}
		}
		art += "\n"
	}
	fmt.Println(art)
}