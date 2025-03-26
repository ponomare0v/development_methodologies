package games

import (
	"math/rand"
	"strconv"
	"time"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

type lcmGame struct {
	question []int
	answer   int
	rules    string
}

func (cur *lcmGame) Generate() error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Используем локальный генератор
	cur.question = make([]int, 3)
	for i := range cur.question {
		cur.question[i] = rng.Intn(50) + 1
	}
	cur.answer = lcm(cur.question[0], lcm(cur.question[1], cur.question[2]))
	return nil
}

func (cur *lcmGame) GetAnswer() int {
	return cur.answer
}

func (cur *lcmGame) GetQuestion() []string {
	question := make([]string, 3)
	for i := 0; i < 3; i++ {
		question[i] = strconv.Itoa(cur.question[i])
	}
	return question
}

func (cur *lcmGame) GetRules() string {
	return cur.rules
}

func LcmGame() Game {
	return &lcmGame{rules: "Find the smallest common multiple of given numbers."}
}
