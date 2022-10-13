package hangman

import (
	"math/rand"
	"time"
)

func RandomLetter(word string) []int {
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
func UncompletedWord(word string) string {
	var underscore string
	for i := 0; i < len(word); i++ {
		underscore += "_"
	}
	wordRune := []rune(word)
	manip := []rune(underscore)
	index := RandomLetter(word)
	for i := 0; i < len(index); i++ {
		manip[index[i]] = wordRune[index[i]]
	}
	return string(manip)
}
