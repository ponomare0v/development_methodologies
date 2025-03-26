package games

type Game interface {
	Generate() error
	GetAnswer() int
	GetQuestion() []string
	GetRules() string
}
