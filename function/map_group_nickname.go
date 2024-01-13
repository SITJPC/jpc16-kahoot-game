package function

import (
	"fmt"
	"kahoot/types/payload"
	"kahoot/utils/value"
	"log"
	"strconv"
	"strings"
)

func MapGroupNickname(str *string, totalScore *int64) *payload.ScoreAdd {
	TeamNoNickname := strings.Split(*str, "-")
	fmt.Println("teamNoNickname: ", TeamNoNickname)
	TeamNo, err := strconv.Atoi(TeamNoNickname[0])
	if err != nil {
		log.Println("Unable to convert groupNo to int", err)
	}

	addScore := &payload.ScoreAdd{
		TeamNo: value.Ptr(int64(TeamNo)),
		Score:  totalScore,
	}

	return addScore
}
