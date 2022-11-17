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
	for attempts > 1 { //While the user have more than 1 attempts
		choosenLetter := LetterChoose() //call the function LetterChoose
		fmt.Println()

		wordSaid = AlreadySaid(choosenLetter, wordSaid) //list of the word already choose
		choosenLetter = strings.Replace(choosenLetter, "\n", "", -1)
		wordInProgresse = Reveal(wordToFind, wordInProgresse, choosenLetter) 
		fmt.Println(AsciiArt(wordInProgresse)) //reveal the letter choose if in the word
		fmt.Println()

		if !IsSaid(wordSaid, choosenLetter) {

			if IsPresent(wordToFind, choosenLetter) {
				Position(attempts) //print the hangman
				fmt.Println(string(colorGreen), "__________________________________________", string(colorReset))//graphic feature

			} else {
				attempts--
				Position(attempts)
				fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
			}
			wordSaid = wordSaid + choosenLetter + "\n"
		} else {
			fmt.Println(string(colorRed), "You already choose this letter", string(colorReset))
			attempts--
			Position(attempts)
			fmt.Println(string(colorRed), "__________________________________________", string(colorReset))
		}

		fmt.Println()
		fmt.Println("remaining try :", attempts-1)
		fmt.Println("\n\n") //graphic feature
		if wordInProgresse == wordToFind { //allows to propose a whole word
			return WinOrLoose(attempts, wordToFind)
		}
	}
	return WinOrLoose(attempts, wordToFind)
}

func Reveal(wordToFind string, wordInProgresse string, choosenLetter string) string {
	word := []rune(wordInProgresse)
	index := 0
	if len(choosenLetter) > 1 { //check if choosenLetter is a word or a letter
		for ii, valuechoosenLetter := range choosenLetter {
			for jj, valueWordToFind := range wordToFind {
				if ii == jj && valueWordToFind != valuechoosenLetter { //check each letter of choosenLetter and wordToFind to know if is egals
					return wordInProgresse
				}
			}
		}
		return wordToFind
	} else {
		for _, letter := range wordToFind {
			for _, valueLettreChoose := range choosenLetter {
				if string(letter) == string(valueLettreChoose) { //check if choosenLetter is in the wordToFind
					word[index] = rune(letter)
				}
			}
			index++
		}
	}
	return string(word)
}

func Position(attempts int) string{
	hangFile, _ := ioutil.ReadFile("../position_hangman.txt") 
	file := strings.Split(string(hangFile), ",,") //crete a table of string with the position of the hangman
	position := 11 - attempts //relates the position to the attemps
	if position < 0 {
		fmt.Println(file[0])
		return file[0]
	} else {
		fmt.Println(file[position])
		return file[position]
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
	file := strings.Split(string(fileIncome), ",,") //crete a table of string with each Ascii letter
	var letter []string
	word := []rune(wordUncomplet)
	var art string
	for j := 0; j < 8; j++ { //loop to define the layer of the ascii letter to print
		for k := 0; k < len(word); k++ { //loop to define all the letter to transform
			if word[k] >= 'a' && word[k] <= 'z' { 
				difference := int(word[k]) - int('a')
				letter = strings.Split(string(file[33+difference]), "\n")
				art += strings.Replace(letter[j], "\r", "", -1)
			} else { // if it's not a letter print the underscore
				letter = strings.Split(string(file[63]), "\n")
				art += strings.Replace(letter[j], "\r", "", -1)
			}
		}
		art += "\n"
	}
	return art
}

func ResultDisplay(word string, firstOutcome string) {
	result := VerifeChar(word, firstOutcome) //varaiable for the answer
	duration, _ := time.ParseDuration("150ms") //define a duration off sleep for the program
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run() // clear the terminal

	occ := 0
	var color string

	for i := 31; occ < 40; i++ { //loop to define the color of the string
		if i == 37 {
			i = 31
		}
		color = strconv.Itoa(i)

		fmt.Print(cmdClear)
		fmt.Println("\033[1;" + color + "m")
		fmt.Println(result)
		if !conclusion { // if the user loose it print the hangman as well
			Position(1)
		}

		time.Sleep(duration)

		occ++
	}
	fmt.Print(colorReset)
}
