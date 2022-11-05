package hangman

import (
	"math/rand"
	"strings"
	"time"
	"os"
	"fmt"
	"bufio"
)

func ChooseWord(fileIncome []byte) string {
	file := strings.Split(string(fileIncome), "\n") //split each word in a string table
	rand.Seed(time.Now().UnixNano())
	randWord := rand.Intn(len(file)) //choose a random index of the table
	wordRune := []rune(file[randWord])
	var word []rune
	var wordString string
	for i := 0; i < len(wordRune); i++ { //this loop prevent error from special character
		if wordRune[i] >= 'a' && wordRune[i] <= 'z' {
			word = append(word, wordRune[i])
		}
		if wordRune[i] == 'é' || wordRune[i] == 'è' {
			word = append(word, 'e')
		}
		if wordRune[i] == 'ç' {
			word = append(word, 'c')
		}
	}
	for _, letter := range word { //convert the rune word into a string
		wordString += string(letter)
	}
	return wordString
}

func AddWord (fileIncome []byte) {
	isValid := true
	list := strings.Split(string(fileIncome), "\n") //split each word in a string table
	if conclusion {
		fmt.Print("Do you want to add a word to the list ? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		add,_ := reader.ReadString('\n')
		add = strings.Replace(add,"\n","",-1)
		
		if add == "yes" {
			f,_ := os.OpenFile("../words.txt", os.O_APPEND|os.O_WRONLY, 0644) //open the file to read and write

			fmt.Print("choose a word : ")
			reader := bufio.NewReader(os.Stdin)
			word,_ := reader.ReadString('\n')
			word = strings.Replace(word,"\n","",-1)


			char := []rune(word)
			
			for j := 0; j < len(char); j++ { //this loop prevent add of invalid character
				if (char[j] != 'é' && char[j] != 'è') && (char[j] < 'a' || char[j] > 'z') {
					isValid = false
					fmt.Println("You entered an invalid character")
				}
			}
			
			if isValid {
				for i := 0; i < len(list); i++ { // verif if the word is already in the list
					if word == list[i] {
						fmt.Println("This word is already in the list")
						isValid = false
						break
					} else {
						isValid = true
					}
				}
			}
			if isValid {
				fmt.Fprint(f,"\n")
				fmt.Fprint(f,word) //add the word to the list
			} else {
				AddWord(fileIncome)
			}
		}
	}
}