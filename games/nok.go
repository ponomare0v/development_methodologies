package games

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Функция для нахождения НОК
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmOfThree(a, b, c int) int {
	return lcm(lcm(a, b), c)
}

func PlayNokGame() {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("Find the smallest common multiple of given numbers.")

	questions := [][]int{
		{5, 7, 15},
		{100, 50, 1},
		{3, 9, 27},
	}

	for _, nums := range questions {
		fmt.Printf("Question: %d %d %d\n", nums[0], nums[1], nums[2])
		fmt.Print("Your answer: ")
		scanner.Scan()
		answer, _ := strconv.Atoi(scanner.Text())

		correctAnswer := lcmOfThree(nums[0], nums[1], nums[2])
		if answer == correctAnswer {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", answer, correctAnswer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}
	}

	fmt.Printf("Congratulations, %s!\n", name)
}
