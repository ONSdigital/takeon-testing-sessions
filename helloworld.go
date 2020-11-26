package main //1

import (
	"errors"
	"fmt"
	"math/rand"
)

//import "os"

type location struct {
	floor         int
	ladderPresent bool
}

//Initialise location (floor=4,ladder=false)		Location object?

//loop forever:
//	Determine location information				Random element (25% of giving ladder=true)
//	Print location information
//	User makes choice							Capture keyboard input
//	Move to new location						climb ladder (U) or move (M)
//	If location is floor 0:						Exit option
//		Print victory message
//		exit game

func main() { //3

	var currentLocation location = location{floor: 4, ladderPresent: false}

	var gameOver = false

	// loop forever:
	for !gameOver {

		//	Determine location information				Random element (25% of giving ladder=true)
		currentLocation.ladderPresent = isLadderPresent()

		printLocationDetails(&currentLocation)

		//	User makes choice							Capture keyboard input
		userInput, inputError := processUserInput()
		if inputError != nil {
			continue
		}

		//	Move to new location						climb ladder (U) or move (M)
		movePlayer(userInput, &currentLocation)

		gameOver = isGameOver(currentLocation.floor)
	}
	// Victory message
	fmt.Println("You have won!")
}

func printLocationDetails(Location *location) {
	fmt.Printf("You are currently on floor : %d\n", Location.floor)
	fmt.Printf("Ladder: %t\n", Location.ladderPresent)
	fmt.Printf("Please enter ('U' for up or 'M' for move):\n")
}

func processUserInput() (string, error) {
	userInput := getInput()
	if !isInputValid(userInput) {
		fmt.Printf("Please enter a valid input \n")
		return "", errors.New("Invalid user input")
	}
	return userInput, nil
}

// To do: distinguish between m&u and what movePlayer function will return
func movePlayer(userInput string, currentLocation *location) {
	if currentLocation.ladderPresent && userInput == "u" || userInput == "U" {
		currentLocation.floor--
	}
}

func isLadderPresent() bool {
	randomNumber := rand.Intn(4)
	if randomNumber == 0 {
		return true
	}
	return false
}

func isInputValid(input string) bool {
	if input == "M" || input == "m" {
		return true
	}
	if input == "U" || input == "u" {
		return true
	}
	return false
}

func getInput() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func isGameOver(floorNumber int) bool {
	if floorNumber == 0 {
		return true
	}
	return false
}

// MODULUS for next time

// if currentPlayer.currentFloor == 0 {
// 	gameOver = true
// 	fmt.Println("Damn you Win again")

// 	// fmt.Print("You can move (M/m)")

// 	// var input string
// 	// e, err := fmt.Scanln(&input)
// 	// if err != nil {
// 	// 	fmt.Fprintln(os.Stderr, err)
// 	// 	return
// 	// }
// }

//To do: Mapping to pseudo code
