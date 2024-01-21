package mongodb

import (
	"fmt"
	"m165/nini/mongodb_quiz/internal/question"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

const pokemonAmount int = 1

func GetAll() {
    execute(bson.D{}, bson.D{})
}

func GetPokemon(q question.Question) {
    proj := bson.D{{"_id", 0}, {"name", 1}, {q.GetWhatValue(), 1}}
    query := bson.D{{"id", q.GetPokemonId()}}
    execute(query, proj)
}

func GetRandomPokemon(amount int, q question.Question) {
    proj := bson.D{{"_id", 0}, {"name", 1}, {q.GetWhatValue(), 1}, {"id", 1}}
    query := bson.D{{"id", randomNum()}}
    execute(query, proj)
}

func randomNum() int{
    min := 1
    max := pokemonAmount
    num := rand.Intn(max - min) + min
    fmt.Println(num)
    return num
}
