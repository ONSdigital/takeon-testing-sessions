package main //1

import "fmt" //2
import "os"

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
		
		fmt.Print("Pick a Floor : ")
		var input int
		e, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		currentPlayer.currentFloor = e
		if currentPlayer.currentFloor == 0 {
			gameOver=true
			fmt.Println("Damn you Win again")
		}
	} 

}
