package games

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func PlayProgressionGame() {
	fmt.Println("Welcome to the Brain Games!")
	fmt.Print("May I have your name? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("What number is missing in the progression?")

	// Создаём локальный генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	questions := generateProgressions(r, 3)

	for _, q := range questions {
		fmt.Printf("Question: %s\n", q.question)
		fmt.Print("Your answer: ")
		scanner.Scan()
		answer, _ := strconv.Atoi(scanner.Text())

		if answer == q.correctAnswer {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", answer, q.correctAnswer)
			fmt.Printf("Let's try again, %s!\n", name)
			return
		}
	}

	fmt.Printf("Congratulations, %s!\n", name)
}

type progressionQuestion struct {
	question      string
	correctAnswer int
}

func generateProgressions(r *rand.Rand, count int) []progressionQuestion {
	var questions []progressionQuestion

	for i := 0; i < count; i++ {
		start := r.Intn(5) + 1
		ratio := r.Intn(5) + 2
		length := r.Intn(6) + 5 // Длина от 5 до 10
		missingIndex := r.Intn(length)

		progression := make([]int, length)
		progression[0] = start
		for j := 1; j < length; j++ {
			progression[j] = progression[j-1] * ratio
		}

		correctAnswer := progression[missingIndex]
		progression[missingIndex] = -1

		questionStr := formatProgression(progression)
		questions = append(questions, progressionQuestion{questionStr, correctAnswer})
	}

	return questions
}

func formatProgression(prog []int) string {
	result := ""
	for _, num := range prog {
		if num == -1 {
			result += ".. "
		} else {
			result += fmt.Sprintf("%d ", num)
		}
	}
	return result[:len(result)-1]
}
