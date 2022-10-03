package hangman

import (
	"bufio"
	"os"
)

func LettreChoose() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
