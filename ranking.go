package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// Player represents a player's record in the ranking
type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

const rankingFile = "ranking.json"

// loadRanking loads players from ranking.json file
func loadRanking() []Player {
	var players []Player

	// Check if file exists
	if _, err := os.Stat(rankingFile); os.IsNotExist(err) {
		return players
	}

	data, err := ioutil.ReadFile(rankingFile)
	if err != nil {
		fmt.Println("Error reading ranking file:", err)
		return players
	}

	err = json.Unmarshal(data, &players)
	if err != nil {
		fmt.Println("Error parsing ranking file:", err)
		return players
	}

	return players
}

// saveRanking saves players to ranking.json file
func saveRanking(players []Player) {
	data, err := json.MarshalIndent(players, "", "  ")
	if err != nil {
		fmt.Println("Error encoding ranking:", err)
		return
	}

	err = ioutil.WriteFile(rankingFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing ranking file:", err)
		return
	}
}

// nameExists checks if a player name already exists in the ranking
func nameExists(players []Player, name string) bool {
	for _, player := range players {
		if player.Name == name {
			return true
		}
	}
	return false
}

// showRanking displays the top players sorted by time (ascending)
func showRanking() {
	players := loadRanking()

	if len(players) == 0 {
		fmt.Println("\nNo players in ranking yet!")
		return
	}

	// Sort by time (ascending - fastest first)
	sort.Slice(players, func(i, j int) bool {
		return players[i].Time < players[j].Time
	})

	fmt.Println("\n=== TOP 10 RANKING ===")
	fmt.Println("Rank | Name              | Time (seconds)")
	fmt.Println("-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-")

	limit := len(players)
	if limit > 10 {
		limit = 10
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%4d | %-17s | %d\n", i+1, players[i].Name, players[i].Time)
	}
}
