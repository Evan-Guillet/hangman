package hangman

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const colorRed = "\033[1;31m"
const colorGreen = "\033[1;32m"
const colorReset = "\033[0m"
const cmdClear = "\033[H\033[2J"

var conclusion bool

func LetterChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choose a letter :")
	reader := bufio.NewReader(os.Stdin)    // create a stdin to get what player enter
	text, error := reader.ReadString('\n') // the char '\n' break the stdin and returne in text what player said

	if error == nil { //check if there is no error
		if text == "é\n" || text == "è\n" { //replace specials charachtere by is simple forme
			text = "e\n"
		}
		if text == "ç\n" {
			text = "c\n"
		}
	}
	return text
}
func IsPresent(wordToFind string, choosenLetter string) bool { // func returne true if letter choose by user is present in word to find
	if len(choosenLetter) > 1 { // check if choosenLetter is a letter or a word
		for ii, choosenLetterRune := range choosenLetter {
			for jj, wordToFindRune := range wordToFind {
				if ii == jj && wordToFindRune != choosenLetterRune { //check each letter of choosenLetter and wordToFind to know if is egals
					return false
				}
			}
		}
		return true
	} else {
		for _, wordRune := range wordToFind {
			for _, choosenLetterRune := range choosenLetter {
				if string(wordRune) == string(choosenLetterRune) { //check if choosenLetter is in the wordToFind
					return true
				}
			}
		}
	}
	return false
}

func AlreadySaid(choosenLetter string, wordSaid string) string {
	err := false
	var said []string
	letter := strings.Replace(choosenLetter, "\n", "", -1)
	if letter < "a" || letter > "z" { // check if all char are letter and returne if its true or false
		err = true
	}

	if err {
		fmt.Println(string(colorRed), "You entered an invalid letter", string(colorReset)) // if there is other char than  letters print err message
	}

	said = strings.Split(wordSaid, "\n") //split by char '\n'

	wordString := strings.Join(said, " ")
	fmt.Println("Already tried :", wordString)

	return wordSaid
}

func IsSaid(wordSaid string, choosenLetter string) bool {
	isSaid := false
	said := []rune(wordSaid)             // creat rune table of wordSaid
	for i := 0; i < len(wordSaid); i++ { // for loop of the lenght of wordSaid
		isSaid = false
		if !isSaid {

			if choosenLetter == string(said[i]) { // if choosenLtter is already said returne true else return false
				isSaid = true
				break
			}
		}
	}
	return isSaid
}

func VerifeChar(wordToFind string, wordUncomplet string) string {
	attempts := 11 // number of attempts + 1 for the start
	var wordSaid string
	wordInProgresse := wordUncomplet
	fmt.Println(AsciiArt(wordUncomplet))
	for attempts > 1 { //
		letterChoose := LetterChoose()
		fmt.Println()

		wordSaid = AlreadySaid(letterChoose, wordSaid)
		letterChoose = strings.Replace(letterChoose, "\n", "", -1)
		wordInProgresse = Reveal(wordToFind, wordInProgresse, letterChoose)
		fmt.Println(AsciiArt(wordInProgresse))
		fmt.Println()

		if !IsSaid(wordSaid, letterChoose) {

			if IsPresent(wordToFind, letterChoose) {
				Position(attempts)
				fmt.Println(string(colorGreen), "__________________________________________", string(colorReset))

			} else {
				attempts--
				Position(attempts)
				fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
			}
			wordSaid = wordSaid + letterChoose + "\n"
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
	if len(letterChoose) > 1 { //check if choosenLetter is a word or a letter
		for ii, valueLetterChoose := range letterChoose {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valueLetterChoose { //check each letter of choosenLetter and wordToFind to know if is egals
					return wordInProgresse
				}
			}
		}
		return wordToFind
	} else {
		for _, letter := range wordToFind {
			for _, valueLettreChoose := range letterChoose {
				if string(letter) == string(valueLettreChoose) { //check if choosenLetter is in the wordToFind
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
	position := 11 - attempts //
	if position < 0 {
		fmt.Println(file[0])
	} else {
		fmt.Println(file[position])
	}
}

func WinOrLoose(attempts int, wordToFind string) string {
	var endPrint string
	if attempts == 1 { // if the player as no more attempts print loose message of if he win print win message
		endPrint = "Dommage ! Vous avez perdu, le mot était :" + AsciiArt(wordToFind)
		conclusion = false
	} else {
		endPrint = "Bravo ! Vous avez gagné, le mot était :" + AsciiArt(wordToFind)
		conclusion = true

	}
	return endPrint
}

func AsciiArt(wordUncomplet string) string {
	fileIncome, _ := ioutil.ReadFile("../standard.txt")
	file := strings.Split(string(fileIncome), ",,")
	var letter []string
	word := []rune(wordUncomplet)
	var art string
	for j := 0; j < 8; j++ {
		for k := 0; k < len(word); k++ {
			if word[k] >= 'a' && word[k] <= 'z' {
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
	return art
}

func ResultDisplay(word string, firstOutcome string) {
	result := VerifeChar(word, firstOutcome)
	duration, _ := time.ParseDuration("150ms")
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	occ := 0
	var color string

	for i := 31; occ < 40; i++ {
		if i == 37 {
			i = 31
		}
		color = strconv.Itoa(i)

		fmt.Print(cmdClear)
		fmt.Println("\033[1;" + color + "m")
		fmt.Println(result)
		if !conclusion {
			Position(1)
		}

		time.Sleep(duration)

		occ++
	}
	fmt.Print(colorReset)
}
