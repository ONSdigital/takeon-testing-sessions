package main //1

import (
		"fmt" 
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
	var inputNumber int = 1

	for gameOver == false {
		fmt.Printf("You are currently on floor : %d\n", currentPlayer.currentFloor)
		output := isLadderPresent(inputNumber)
		fmt.Printf("Ladder: %t\n", output)
		if output == true {
			currentPlayer.currentFloor--
		}
		gameOver = isGameOver(currentPlayer.currentFloor)
		inputNumber++
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
	return ""
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

// To do: need to link new functions (isInputValid etc) into main loop