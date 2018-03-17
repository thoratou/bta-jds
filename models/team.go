package models

//Team team data
type Team struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	ManagerID string                `json:"managerid"`
	Players   []*PlayerData         `json:"players"`
	GameID    string                `json:"gameid"`
	Comment   string                `json:"comment"`
	Forum     map[string]*ForumPost `json:"forum"`
	Removed   bool                  `json:"removed"`
}

//Teams team list
type Teams struct {
	Map map[string]*Team `json:"teams"`
}
