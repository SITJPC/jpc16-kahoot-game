package function

import "kahoot/utils/config"

func ChooseToken(gameType string) *string {
	if gameType == "movie" {
		return config.C.MovieToken
	} else {
		return config.C.SongToken
	}
}
