package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Hand representa una jugada posible y contra qué gana
type Hand struct {
	Hand        string
	WinsAgainst [2]string
}

var (
    rock     = Hand{"Rock", [2]string{"Scissors", "Lizard"}}
    paper    = Hand{"Paper", [2]string{"Rock", "Spock"}}
    scissors = Hand{"Scissors", [2]string{"Paper", "Lizard"}}
    lizard   = Hand{"Lizard", [2]string{"Spock", "Paper"}}
    spock    = Hand{"Spock", [2]string{"Scissors", "Rock"}}

    handOptions = [5]Hand{rock, paper, scissors, lizard, spock}
)


func victory() string {
	return "Victory :)"
}

func draw() string {
	return "Draw"
}

func defeat() string {
	return "Defeat :("
}

func addEmoji(hand string) string {
    switch hand {
    case "Rock":
        return "Rock 🪨"
    case "Paper":
        return "Paper 📄"
    case "Scissors":
        return "Scissors ✂️"
    case "Lizard":
        return "Lizard 🦎"
    case "Spock":
        return "Spock 🖖"
    default:
        return hand
    }
}

func getValidPlayerInput() string {
	var playerInput string
	validOptions := []string{"rock", "paper", "scissors", "lizard", "spock"}

	for {
		fmt.Print("Enter your hand (Rock, Paper, Scissors, Lizard, Spock): ")
		fmt.Scanln(&playerInput)
		playerInput = strings.ToLower(strings.TrimSpace(playerInput))

		for _, option := range validOptions {
			if playerInput == option {
				return playerInput
			}
		}
		fmt.Println("Invalid input. Please choose Rock, Paper, Scissors, Lizard, or Spock.")
	}
}

func randomizeComputerHand(handOptions [5]Hand) Hand {
	randomGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randomGen.Intn(len(handOptions))
	return handOptions[randomIndex]
}

func getPlayerHand(playerInput string, handOptions [5]Hand) (Hand, error) {
	playerInput = strings.ToLower(strings.TrimSpace(playerInput))
	
	for _, hand := range handOptions {
		// Comparar solo el texto sin emojis
		handName := strings.ToLower(strings.TrimSpace(hand.Hand))
		if handName == playerInput {
			return hand, nil
		}
	}
	
	return Hand{}, fmt.Errorf("invalid choice, please try again with rock, paper, scissors, lizard or spock")
}

func playGame(playerHand, computerHand Hand) string {
	if playerHand.Hand == computerHand.Hand {
		return draw()
	}

	for _, winningHand := range playerHand.WinsAgainst {
		if winningHand == computerHand.Hand {
			return victory()
		}
	}

	return defeat()
}

func main() {
	fmt.Println("\n🎮 Welcome to Rock🪨, Paper📄, Scissors✂️, Lizard🦎, Spock🖖! 🎮")
	fmt.Println("Ready? :D")
    time.Sleep(2 * time.Second) 
	fmt.Println("Set...")
    time.Sleep(2 * time.Second) 
	fmt.Println("3")
    time.Sleep(1 * time.Second) 
	fmt.Println("2")
    time.Sleep(1 * time.Second) 
	fmt.Println("1")
    time.Sleep(1 * time.Second) 
	fmt.Println("Go!")
	time.Sleep(1 * time.Second) 

	// Obtener una entrada válida del jugador
	playerInput := getValidPlayerInput()

	// Seleccionar una mano aleatoria para la computadora
	computerHand := randomizeComputerHand(handOptions)

	// Obtener la mano del jugador
	playerHand, err := getPlayerHand(playerInput, handOptions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

// Mostrar las elecciones con emojis
fmt.Printf("\nPlayer chose: %s\n", addEmoji(playerHand.Hand))
fmt.Printf("Computer chose: %s\n", addEmoji(computerHand.Hand))

result := playGame(playerHand, computerHand)
fmt.Printf("\nResult: %s\n\n", result)
}