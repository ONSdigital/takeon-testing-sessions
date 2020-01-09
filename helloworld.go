package main //1

import "fmt" //2

type player struct{
	health int
	name string
	currentFloor int
}


func main() { //3
	var startingFloor int = 4
	var gameOver = false
	var currentPlayer player  
	currentPlayer.currentFloor = startingFloor 
	
	for gameOver==false{
		
		fmt.Printf("Current floor : %d\n", currentPlayer.currentFloor)
		currentPlayer.currentFloor--
		if currentPlayer.currentFloor == 0 {
			gameOver=true
			fmt.Println("Damn you Win again")
		}
	} 

}
