package function

import (
	"kahoot/types/payload"
	"kahoot/utils/value"
	"log"
	"strconv"
	"strings"
)

func MapGroupNickname(str *string, totalScore *int64) *payload.ScoreAdd {
	GroupNoNickname := strings.Split(*str, "-")
	groupNo, err := strconv.Atoi(GroupNoNickname[0])
	if err != nil {
		log.Println("Unable to convert groupNo to int", err)
	}
	addScore := &payload.ScoreAdd{
		GroupNo:  value.Ptr(int64(groupNo)),
		Nickname: value.Ptr(GroupNoNickname[1]),
		Score:    totalScore,
	}

	return addScore
}
