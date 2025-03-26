package main

import (
	"fmt"

	"github.com/ponomare0v/development_methodologies/games"
)

func main() {
	fmt.Println("Welcome to the Brain Games!\nMay I have your name? ")
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Hello " + name)
	for {
		fmt.Println("Which game would you like to play?")
		fmt.Println("1 - LCM")
		fmt.Println("2 - Progression")
		fmt.Println("-1 - Exit")

		var choice string
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		var game games.Game
		switch choice {
		case "1":
			game = games.LcmGame()
		case "2":
			game = games.NewProgressionGame()
		case "-1":
			fmt.Println("Goodbye " + name + "!")
			return
		default:
			fmt.Println("Invalid input.")
			continue
		}

		fmt.Println(game.GetRules())
		if err := game.Generate(); err != nil {
			fmt.Println("Error generating game:", err)
			continue
		}

		fmt.Print("Question: ")
		for _, v := range game.GetQuestion() {
			fmt.Print(v, " ")
		}
		fmt.Println()

		fmt.Println("Your answer:")
		var userAns int
		if _, err := fmt.Scan(&userAns); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if userAns == game.GetAnswer() {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", userAns, game.GetAnswer())
			fmt.Printf("Let's try again, %s!\n", name)
		}
	}
}
