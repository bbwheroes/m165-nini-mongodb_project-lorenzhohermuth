package main

import (
	"m165/nini/mongodb_quiz/pkg/tea"
	"m165/nini/mongodb_quiz/pkg/mongodb"
	"m165/nini/mongodb_quiz/internal/question"
)

func main() {
	mongodb.GetPokemon(question.GenerateQuestion())
	mongodb.GetRandomPokemon(3, question.GenerateQuestion())
	tea.Exec()
}
