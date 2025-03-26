package games

import (
	"math/rand"
	"strconv"
	"time"
)

type progressionGame struct {
	sequence []int
	missing  int
	rules    string
}

func (game *progressionGame) Generate() error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Используем локальный генератор

	length := rng.Intn(6) + 5 // Длина прогрессии от 5 до 10
	game.sequence = make([]int, length)
	ratio := rng.Intn(5) + 2            // Множитель (от 2 до 6)
	game.sequence[0] = rng.Intn(10) + 1 // Первое число в прогрессии

	for i := 1; i < length; i++ {
		game.sequence[i] = game.sequence[i-1] * ratio // Геометрическая прогрессия
	}

	game.missing = rng.Intn(length) // Индекс пропущенного числа
	return nil
}

func (game *progressionGame) GetAnswer() int {
	return game.sequence[game.missing]
}

func (game *progressionGame) GetQuestion() []string {
	questions := make([]string, len(game.sequence))
	for i, num := range game.sequence {
		if i == game.missing {
			questions[i] = ".."
		} else {
			questions[i] = strconv.Itoa(num)
		}
	}
	return questions
}

func (game *progressionGame) GetRules() string {
	return game.rules
}

func NewProgressionGame() Game {
	return &progressionGame{rules: "Find the missing number in the geometric progression."}
}
