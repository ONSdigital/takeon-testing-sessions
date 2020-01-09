package main //1

import "fmt" //2
//import "os"

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
		fmt.Printf("You are currently on floor : %d\n", currentPlayer.currentFloor)
		if currentPlayer.currentFloor == 0 {
			gameOver=true
			fmt.Println("Damn you Win again")

		fmt.Print("You can move (M/m)")
	
		// var input string
		// e, err := fmt.Scanln(&input)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, err)
		// 	return
		// }
	} 

}
}

func determineLadder() bool {
	return false
}

func getInput() {

}