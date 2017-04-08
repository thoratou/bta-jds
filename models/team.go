package models

//Team team data
type Team struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ManagerID string   `json:"managerid"`
	Players   []string `json:"players"`
	GameID    string   `json:"gameid"`
	Removed   bool     `json:"removed"`
}

//Teams team list
type Teams struct {
	Map map[string]*Team `json:"teams"`
}
