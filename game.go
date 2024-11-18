package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hand struct {
    Hand        string
	WinsAgainst [2]string // Array fijo para las opciones que gana
};

func Victory() string {
	return "Victory :)"
};

func Draw() string {
	return "Draw. "
};

func Defeat() string {
	return "Defeat :( "
};

func RandomizeComputerHand(handOptions [5]Hand) Hand {
	// Crear un nuevo generador de números aleatorios
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randomGen.Intn(len(handOptions))
	return handOptions[randomIndex]
};

func main() {
	rock := Hand{"Rock 🪨", [2]string{"Scissors ✂️", "Lizard 🦎"}}
	paper := Hand{"Paper 📄", [2]string{"Rock 🪨", "Spock 🖖"}}
	scissors := Hand{"Scissors ✂️", [2]string{"Paper 📄", "Lizard 🦎"}}
	lizard := Hand{"Lizard 🦎", [2]string{"Spock 🖖", "Paper 📄"}}
	spock := Hand{"Spock 🖖", [2]string{"Scissors ✂️", "Rock 🪨"}}

	// Imprimir ejemplo
	fmt.Println("Rock: ", rock)
	fmt.Println("Paper: ", paper)
	fmt.Println("Scissors: ", scissors)
	fmt.Println("Lizard: ", lizard)
	fmt.Println("Spock: ", spock)

	handOptions := [5]Hand{rock, paper, scissors, lizard, spock}
	fmt.Println("HandOptions: ", handOptions)

	// Seleccionar mano aleatoria para la computadora
	computerHand := RandomizeComputerHand(handOptions)
	fmt.Println("Computer's Hand: ", computerHand)
};



// 	 JS
	
// 	const getPlayerHand = (playerInput) => {
// 		let playerHand = handOptions.find(hand => hand.hand === playerInput);
// 		// console.log("🚀 ~ getPlayerHand ~ playerHand:", playerHand)
// 		return playerHand ? playerHand : "Su elección no es posible, pruebe con piedra, papel o tijera"
// 	};
	
// 	const playGame = (playerInput) => {
// 		let playerHand = getPlayerHand(playerInput)
// 		console.log("🚀 Player Hand:", playerHand)
// 		let computerHand = randomizeComputerHand()
// 		console.log("🚀 Computer Hand:", computerHand)
// 		if (playerHand.hand === computerHand.hand) return draw()
// 		else if (playerHand.winsAgainst === computerHand.hand) return victory()
// 		else return defeat()
// 	};
	
// 	playGame("tijera")
// }
