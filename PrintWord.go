package hangman

import (
	"math/rand"
	"time"
)

func RandomLetter(word string) []int { //func that choose how many letter to print at the start
	n := len(word)/2 - 1
	table := []rune(word)
	var index []int
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(len(table))
		index = append(index, randInt)
	}
	return index
}
func UncompletedWord(word string) string {// func that return the word with n letter revealed
	var underscore string
	for i := 0; i < len(word); i++ { //add as many underscore as number of letter
		underscore += "_"
	}
	wordRune := []rune(word)
	manip := []rune(underscore)
	index := RandomLetter(word)
	for i := 0; i < len(index); i++ { //replace underscore with letter
		manip[index[i]] = wordRune[index[i]]
	}
	return string(manip)
}
