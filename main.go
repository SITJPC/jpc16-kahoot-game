package main

import (
	"kahoot/function"
	"os"
	"path/filepath"
)

func main() {
	// * Receive 2 argument
	gameType := os.Args[1]
	file := os.Args[2]

	// * Initial token
	token := function.ChooseToken(gameType)

	// * Join path
	filePath := filepath.Join("resources", gameType, file)

	// * Map struct
	mapScores := function.OpenAndGetDataFromExcel(filePath)

	for _, score := range mapScores {
		addScore := function.MapGroupNickname(score.Player, score.TotalScore)
		function.DoReq(addScore, token)
	}

}
