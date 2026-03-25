# Math Game in Go

A simple command-line game where the player solves math problems as fast as possible.

## Features

- **Interactive Menu** - Navigate between playing, viewing rankings, and exiting
- **Timed Challenges** - Race against the clock to solve math problems
- **Player Ranking** - Persistent leaderboard saved to JSON
- **Unique Player Names** - Each player is tracked individually
- **Score Calculation** - Points based on speed and accuracy

## Prerequisites

- Go 1.13 or higher
- Terminal/Command line access

## Build Instructions

### Compile the Game

```bash
cd /path/to/math_game
go build -o math-game
```

This creates an executable named `math-game` in your project directory.

Alternatively, build and run directly:

```bash
go run main.go game.go ranking.go
```

## How to Run

### Run the compiled executable:

```bash
./math-game
```

### On Windows:

```bash
math-game.exe
```

## How to Play

1. **Start the program** - Run `./math-game`
2. **Choose an option** from the menu:
   - **Option 1**: Start a game
   - **Option 2**: View the ranking leaderboard
   - **Option 3**: Exit the program

3. **Start a Game**:
   - Enter your player name (must be unique)
   - Solve math addition problems as quickly as possible
   - Each correct answer gives you points
   - Reach 100 points to complete the game

4. **Your Results**:
   - After finishing, your total time is recorded
   - Your result is automatically saved to the ranking
   - View all rankings from the main menu

## Game Mechanics

- **Questions**: Simple addition problems (0-100)
- **Scoring**: 20 points per correct answer
- **Win Condition**: Reach 100 points (5 correct answers)
- **Ranking**: Sorted by fastest completion time

## File Structure

```
math_game/
├── main.go          # Menu entry point
├── game.go          # Game logic and game loop
├── ranking.go       # Player ranking and JSON persistence
├── game_test.go     # Unit tests
├── ranking.json     # Player leaderboard (auto-created)
├── go.mod           # Go module file
└── README.md        # This file
```

## Data Storage

Player rankings are automatically saved to `ranking.json` in the format:

```json
[
  {
    "id": 1,
    "name": "Player1",
    "time": 45
  },
  {
    "id": 2,
    "name": "Player2",
    "time": 52
  }
]
```

## Testing

Run the test suite:

```bash
go test -v
```

Tests simulate multiple players completing the game to verify the core game logic works correctly.

## Example Gameplay

```
=== MENU ===
1. Start a game
2. Ranking
3. Exit
Choose option: 1

Enter your name: Alice

5 + 23 = ?
28
12 + 88 = ?
100
...

--- RESULT ---
Total time: 45s
Total score: 222

Your time saved: 45 seconds!
```