package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	pokeapi "github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func getStats(pokemonID int) (structs.Pokemon, error) {
	pokemon, err := pokeapi.Pokemon(strconv.Itoa(pokemonID))

	if err != nil {
		return structs.Pokemon{}, err
	}

	return pokemon, nil

}

func simulateBattle() {

	rand.Seed(time.Now().UnixNano())
	pokemon1ID := rand.Intn(151) + 1
	pokemon2ID := rand.Intn(151) + 1

	pokemon1, _ := getStats(pokemon1ID)
	pokemon2, _ := getStats(pokemon2ID)

	pokemon1Power := pokemon1.Stats[0].BaseStat + pokemon2.Stats[0].BaseStat
	pokemon2Power := pokemon2.Stats[0].BaseStat + pokemon1.Stats[0].BaseStat

	fmt.Println("Battle between", pokemon1.Name, "and", pokemon2.Name)

	if pokemon1Power > pokemon2Power {
		fmt.Println(pokemon1.Name, "wins!")
	} else if pokemon2Power > pokemon1Power {
		fmt.Println(pokemon2.Name, "wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}

func main() {
	simulateBattle()
}
