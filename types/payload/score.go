package payload

import "time"

type Score struct {
	Date            *time.Time `json:"date"`
	Rank            *int       `json:"rank"`
	Player          *string    `json:"player"`
	TotalScore      *int64     `json:"totalScore"`
	CorrectAnswer   *int       `json:"correctAnswer"`
	IncorrectAnswer *int       `json:"incorrectAnswer"`
}

type ScoreAdd struct {
	GroupNo  *int64  `json:"groupNo"`
	Nickname *string `json:"nickname"`
	Score    *int64  `json:"score" `
}
