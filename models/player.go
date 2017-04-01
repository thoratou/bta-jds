package models

//Player player data
type Player struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
}

//Players player list
type Players struct {
	Map map[string]*Player `json:"players"`
}

//PlayerData player data for one game or team
type PlayerData struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
}
