package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello gamer!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Want to start a game? (y/n): ")
	answer, _ := reader.ReadString('\n')

	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer != "y" {
		fmt.Println("Goodbye!")
		return
	}

	game := NewGame()
	game.Start() // старт таймера

	// Основний ігровий цикл
	for !game.IsFinished() {
		question, correct := game.GenerateQuestion()

		fmt.Println(question)

		input, _ := reader.ReadString('\n')

		isCorrect := game.CheckAnswer(input, correct)

		game.Update(isCorrect)
	}

	game.PrintResult()
}
