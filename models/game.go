package models

//Game game data
type Game struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	TeamGame    bool     `json:"teamgame"`
	MinPlayers  int      `json:"minPlayers"`
	Description []string `json:"description"`
}

//Games game list
type Games struct {
	Map map[string]*Game `json:"games"`
}
