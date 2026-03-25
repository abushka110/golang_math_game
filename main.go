package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Start a game")
		fmt.Println("2. Ranking")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Choose option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			startGame()
		case 2:
			showRanking()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
