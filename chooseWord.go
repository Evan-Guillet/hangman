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
    return file[randWord]
}
