package main //1

import "fmt" //2
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
		inputNumber++

	}
}

func isLadderPresent(number int) bool {
	if number == 1 {
		return true
	}
	return false
}

func getInput() {

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
