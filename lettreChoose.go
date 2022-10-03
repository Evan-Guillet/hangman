package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func LettreChoose() string { //func that return a string contane what user write in terminal
	fmt.Print("choisi une lettre :")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
