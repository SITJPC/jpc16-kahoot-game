package main

import (
	"kahoot/function"
	"kahoot/types/payload"
	"kahoot/types/response"
	"kahoot/utils/value"
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

	mapData, _ := value.Iterate(mapScores, func(a *payload.Score) (*payload.ScoreAdd, *response.ErrorInstance) {
		addScore := function.MapGroupNickname(a.Player, a.TotalScore)
		return &payload.ScoreAdd{
			TeamNo: addScore.TeamNo,
			Score:  addScore.Score,
		}, nil
	})

	var sumScore *int64
	for _, data := range mapData {
		// Check if sumScore is nil, initialize it with zero
		if sumScore == nil {
			initialScore := int64(0)
			sumScore = &initialScore
		}

		// Dereference the pointer and update the value
		*sumScore += *data.Score
	}

	finalResult := &payload.ScoreAdd{
		TeamNo: mapData[0].TeamNo,
		Score:  sumScore,
	}

	function.DoReq(finalResult, token)

}
