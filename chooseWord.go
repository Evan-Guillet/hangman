package hangman

import (
    "strings"
    "math/rand"
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
    }
    for _, letter := range word {
        wordString += string(letter)
    }
    return wordString
}
