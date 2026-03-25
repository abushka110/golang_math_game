package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Структура гри, яка зберігає стан гравця
type Game struct {
	score          int
	correctAnswers int
	startTime      time.Time
	fiveTime       time.Duration
	fiveReached    bool
}

func NewGame() *Game {
	return &Game{}
}

// Початок гри (фіксація часу старту)
func (g *Game) Start() {
	g.startTime = time.Now()
}

// Генерація математичного прикладу (додавання)
func (g *Game) GenerateQuestion() (string, int) {
	a := rand.Intn(101)
	b := rand.Intn(101)

	return fmt.Sprintf("%d + %d = ?", a, b), a + b
}

// Перевірка відповіді користувача
func (g *Game) CheckAnswer(input string, correct int) bool {
	input = strings.TrimSpace(input)

	answer, err := strconv.Atoi(input)
	if err != nil {
		// якщо введено не число — відповідь вважається неправильною
		return false
	}

	return answer == correct
}

// Оновлення стану гри після відповіді
func (g *Game) Update(isCorrect bool) {
	if isCorrect {
		g.score += 20
		g.correctAnswers++

		// фіксація часу, коли досягнуто 5 правильних відповідей
		if g.correctAnswers == 5 && !g.fiveReached {
			g.fiveTime = time.Since(g.startTime)
			g.fiveReached = true
		}
	}
}

// Перевірка, чи гра завершена
func (g *Game) IsFinished() bool {
	return g.score >= 100
}

// Вивід результату гри
func (g *Game) PrintResult() {
	totalTime := time.Since(g.startTime) // загальний час гри

	seconds := totalTime.Seconds() // переведення часу у секунди

	// захист від ділення на 0
	if seconds == 0 {
		seconds = 0.001
	}

	// формула очок: чим менше час — тим більше очок
	scoreFromTime := int(10000 / seconds)

	fmt.Println("\n--- RESULT ---")
	fmt.Printf("Total time: %v\n", totalTime)
	fmt.Printf("Total score: %d\n", scoreFromTime)
}

// startGame handles the game flow: name input, gameplay, and ranking save
func startGame() {
	players := loadRanking()

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	if nameExists(players, name) {
		fmt.Println("Name already used!")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	start := time.Now()

	// Create and start the game
	game := NewGame()
	game.Start()

	// Main game loop
	for !game.IsFinished() {
		question, correct := game.GenerateQuestion()

		fmt.Println(question)

		input, _ := reader.ReadString('\n')

		isCorrect := game.CheckAnswer(input, correct)

		game.Update(isCorrect)
	}

	game.PrintResult()

	duration := int(time.Since(start).Seconds())

	newPlayer := Player{
		ID:   len(players) + 1,
		Name: name,
		Time: duration,
	}

	players = append(players, newPlayer)
	saveRanking(players)

	fmt.Printf("Your time saved: %d seconds!\n", duration)
}
