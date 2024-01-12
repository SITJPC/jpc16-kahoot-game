package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"kahoot/types/payload"
	"log"
	"net/http"
)

func DoReq(addScore *payload.ScoreAdd, token *string) {

	addScorePayload, _ := json.Marshal(addScore)
	req, err := http.NewRequest("POST", "http://localhost:3000/api/operate/score/player", bytes.NewBuffer(addScorePayload))
	if err != nil {
		log.Fatal("UNABLE TO CREATE handler-reserve REQUEST")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+*token)

	resq, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("UNABLE TO SEND handler-reserve REQUEST")
	}
	bytes, _ := io.ReadAll(resq.Body)
	spew.Dump(string(bytes))
	if resq.StatusCode == 200 {
		fmt.Printf("GroupNo: %d Nickname: %s Score: %d\n", *addScore.TeamNo, *addScore.Nickname, *addScore.Score)
		log.Println("Successfully add score")
	}

}
