package models

//Player player data
type Player struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Players player list
type Players struct {
	Map map[string]*Player `json:"players"`
}
