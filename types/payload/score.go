package payload

import "time"

type Score struct {
	Date            *time.Time `json:"date"`
	Rank            *int       `json:"rank"`
	Player          *string    `json:"player"`
	TotalScore      *int       `json:"totalScore"`
	CorrectAnswer   *int       `json:"correctAnswer"`
	IncorrectAnswer *int       `json:"incorrectAnswer"`
}
