package main

import (
	"strconv"
	"testing"
	"time"
)

// Симуляція одного користувача
func simulateUser() time.Duration {
	game := NewGame()
	game.Start()

	for !game.IsFinished() {
		_, correct := game.GenerateQuestion()

		answer := strconv.Itoa(correct)

		isCorrect := game.CheckAnswer(answer, correct)

		game.Update(isCorrect)
	}

	return time.Since(game.startTime)
}

// Тест, який симулює 5 користувачів
func TestGameSimulation(t *testing.T) {
	users := 5

	for i := 0; i < users; i++ {
		duration := simulateUser()

		t.Logf("User %d finished in %v", i+1, duration)
	}
}
