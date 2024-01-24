package main

import (
	"fmt"
	"m165/nini/mongodb_quiz/internal/pokemon"
	"m165/nini/mongodb_quiz/internal/question"
	"m165/nini/mongodb_quiz/pkg/mongodb"
	"os"

	"math/rand"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"go.mongodb.org/mongo-driver/bson"
)

const pokemonAmount int = 250
const amountPokemonOnScreen = 4

var points int = 0
var name string = "NO NAME"

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
	}
}

type Model struct {
    Question string
    Answers  []string
    RightIndex []int
    Index   int 
    nameInput textinput.Model
    isStartScreen bool
    isEndScreen bool
}

func initModel() Model {
    ti := textinput.New()
    ti.Placeholder = "Max"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20
    return Model {
	isStartScreen: true,
	nameInput: ti,
    }
}

func (m Model) Init() tea.Cmd {
    return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    if m.isStartScreen {
	return handleStartScreen(cmd, msg, m)
    } else {
	return handleGameScreen(cmd, msg, m)
    }
}

func handleStartScreen(cmd tea.Cmd, msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
    m.nameInput, cmd = m.nameInput.Update(msg)
    switch msg := msg.(type) {
	case tea.KeyMsg:
	    // Cool, what was the actual key pressed?
	    switch msg.String() {
	    case "enter":
		name = m.nameInput.Value()
		return build(question.GenerateQuestion(), amountPokemonOnScreen), nil
	}
    }
    return m, cmd
}

func handleGameScreen(cmd tea.Cmd, msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
	case tea.KeyMsg:

	    // Cool, what was the actual key pressed?
	    switch msg.String() {

	    // These keys should exit the program.
	    case "ctrl+c", "q":
		return m, tea.Quit

	    // The "up" and "k" keys move the cursor up
	    case "up", "k":
		if m.Index > 0 {
		    m.Index--
		}

	    // The "down" and "j" keys move the cursor down
	    case "down", "j":
		if m.Index < len(m.Answers)-1 {
		    m.Index++
		}

	    // The "enter" key and the spacebar (a literal space) toggle
	    // the selected state for the item that the cursor is pointing at.
	    case "enter", " ":
		if contains(m.RightIndex, m.Index) {
		    points++
		}
		return build(question.GenerateQuestion(), amountPokemonOnScreen), nil
	    }
    }
    return m, nil
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func (m Model) View() string {
    var s string
    if m.isStartScreen {
	s += m.nameInput.View()
    } else if m.isEndScreen{

    } else {
	s += fmt.Sprintf("Player: %v\n", name)
	s += fmt.Sprintf("Points: %v\n", points)
	s += m.Question + "\n"
	for i, ans := range m.Answers{
	    cursor := " " // no cursor
	    if m.Index == i {
		cursor = ">" // cursor!
	    }
	    s += fmt.Sprintf("%s %s\n", cursor, ans)
	}
    }
    return s
}

func getRandomPokemon(amount int, q question.Question) []mongodb.MongoPokemon{
    pokemons := []mongodb.MongoPokemon{}
    for i := 0 ; i < amount ; i++ {
	query := bson.D{{"id", randomNum()}}
	pokemons = append(pokemons, mongodb.Execute(query))
    }
    return pokemons
}

func toPokemon(mp []mongodb.MongoPokemon ,q question.Question) []pokemon.Pokemon {
    p := make([]pokemon.Pokemon, len(mp))
    for i, v := range mp {
	p[i] = pokemon.Pokemon{
	    Name: v.Name,
	    Value: v.GetValue(q.GetWhatValue()),
	}
    }
    return p
}

func build(q question.Question, amount int) Model {
	mongoPokemonList := getRandomPokemon(amount, q)
	pokemonList := toPokemon(mongoPokemonList, q)
	answer := pokemonList[rand.Intn(amount)].GetValue()
	var correctIndex []int
	pokemonNameAnswer := make([]string, amount)
	for i, v := range pokemonList {
		pokemonNameAnswer[i] = v.GetName()
		if v.GetValue() == answer {
			correctIndex = append(correctIndex, i)
		}
	}
	return Model {
		Question: q.GetWhatValue() + " is " + answer,
		Answers:  pokemonNameAnswer,
		RightIndex: correctIndex,
		Index: 0, 
	}  
}

func randomNum() int{
    min := 1
    max := pokemonAmount
    num := rand.Intn(max - min) + min
    return num
}
