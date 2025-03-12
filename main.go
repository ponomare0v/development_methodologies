// package main

// import (
// 	nok "github.com/ponomare0v/development_methodologies/games"
// 	progression "github.com/ponomare0v/development_methodologies/games"
// )

// func main() {
// 	nok.PlayNokGame()
// 	progression.PlayProgressionGame()
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	nok "github.com/ponomare0v/development_methodologies/games"
	progression "github.com/ponomare0v/development_methodologies/games"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Brain Games!")
	fmt.Println("Which game would you like to play?")
	fmt.Println("1 - Least Common Multiple (nok)")
	fmt.Println("2 - Geometric Progression (progression)")
	fmt.Print("Enter 1 or 2: ")

	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "1":
		nok.PlayNokGame()
	case "2":
		progression.PlayProgressionGame()
	default:
		fmt.Println("Invalid choice. Please restart and enter 1 or 2.")
	}
}
