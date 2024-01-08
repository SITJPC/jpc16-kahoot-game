package function

import (
	"kahoot/types/payload"
	"kahoot/utils/value"
	"strconv"
	"time"
)

func MapTableToScoreStruct(rows [][]string) ([]*payload.Score, error) {
	scores := make([]*payload.Score, 0)
	for i, row := range rows {
		// * Skip 2 header
		if i > 2 {
			var score payload.Score
			score.Date = value.Ptr(time.Now())
			for j, colCell := range row {
				if colCell != "-" {
					switch j {
					case 1:
						score.Player = value.Ptr(colCell)
					case 0, 2, 3, 4:
						numbers, _ := strconv.Atoi(colCell)
						if j == 0 {
							score.Rank = value.Ptr(numbers)
						} else if j == 2 {
							totalScore, _ := strconv.Atoi(colCell)
							score.TotalScore = value.Ptr(int64(totalScore))
						} else if j == 3 {
							score.CorrectAnswer = value.Ptr(numbers)
						} else {
							score.IncorrectAnswer = value.Ptr(numbers)
						}
					}
				}
			}
			scores = append(scores, &score)
		}
	}
	return scores, nil
}
