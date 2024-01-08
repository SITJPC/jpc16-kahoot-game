package config

type config struct {
	MovieToken *string `json:"movieToken" yaml:"movie_token"`
	SongToken  *string `json:"songToken" yaml:"song_token"`
}
