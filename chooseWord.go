package hangman

import (
	"math/rand"
	"strings"
	"time"
)

func ChooseWord(fileIncome []byte) string {
	file := strings.Split(string(fileIncome), "\n")
	rand.Seed(time.Now().UnixNano())
	randWord := rand.Intn(len(file))
	wordRune := []rune(file[randWord])
	var word []rune
	var wordString string
	for i := 0; i < len(wordRune); i++ {
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
	for _, letter := range word {
		wordString += string(letter)
	}
	return wordString
}
