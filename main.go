package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func getStats(pokemonID int) (structs.Pokemon, error) {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(pokemonID))

	if err != nil {
		return structs.Pokemon{}, err
	}

	return pokemon, nil

}

func sumBaseStats(pokemon structs.Pokemon) int {
	total := 0

	for _, stat := range pokemon.Stats {
		total += stat.BaseStat
	}

	return total
}

func simulateBattle() {

	pokemon1ID := r.Intn(151) + 1
	pokemon2ID := r.Intn(151) + 1

	pokemon1, _ := getStats(pokemon1ID)
	pokemon2, _ := getStats(pokemon2ID)

	winner := determineWinner(pokemon1, pokemon2)

	fmt.Println("Battle between", pokemon1ID, capitalizeName(pokemon1.Name), "and", pokemon2ID, capitalizeName(pokemon2.Name))
	fmt.Println(winner, "wins!")
}

func determineWinner(pokemon1 structs.Pokemon, pokemon2 structs.Pokemon) string {
	pokemon1Power := sumBaseStats(pokemon1)
	pokemon2Power := sumBaseStats(pokemon2)

	if pokemon1Power > pokemon2Power {
		return capitalizeName(pokemon1.Name)
	} else if pokemon2Power > pokemon1Power {
		return capitalizeName(pokemon2.Name)
	} else {
		return "It's a tie!"
	}
}

// Function to capitalize the first letter of the Pokemon's names
func capitalizeName(s string) string {

	if len(s) == 0 {
		return s
	}
	first := strings.ToUpper(string(s[0]))

	if len(s) == 1 {
		return first
	}

	rest := s[1:]
	return first + rest
}

func main() {
	simulateBattle()
}
