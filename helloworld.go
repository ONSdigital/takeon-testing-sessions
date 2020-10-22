package main //1

import (
		"fmt" 
		"errors"
	)
//import "os"



type player struct {
	health       int
	name         string
	currentFloor int
}

func main() { //3
	var startingFloor int = 4
	var gameOver = false
	var currentPlayer player
	currentPlayer.currentFloor = startingFloor
	currentPlayer.health = 20
	var inputNumber int = 1

	for !gameOver{
		fmt.Printf("You are currently on floor : %d\n", currentPlayer.currentFloor)
		ladderPresent := isLadderPresent(inputNumber)
		fmt.Printf("Ladder: %t\n", ladderPresent)
		userInput, inputError := processUserInput()
		if inputError != nil{
			continue
		}
		gameOver = isGameOver(currentPlayer.currentFloor)
		inputNumber++
	}
}

func processUserInput() (string, error) {
	userInput := getInput()
	if !isInputValid(userInput){
		fmt.Printf("Please enter a valid input \n")
		return "", errors.New("Invalid user input")
	}
	return userInput, nil
}

// To do: distinguish between m&u and what movePlayer function will return
func movePlayer(ladderPresent bool, userInput string) {
	if ladderPresent && userInput == "u" || userInput == "U" {
		currentPlayer.currentFloor--
	}
}

func isLadderPresent(number int) bool {
	if number % 4 == 0 {
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

// To do: Simplify main loop, start phase 2 dawn of the monsters, introduce random numbers for ladder/monster