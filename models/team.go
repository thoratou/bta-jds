package models

//Team team data
type Team struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Manager string   `json:"manager"`
	Players []string `json:"players"`
}

//Teams team list
type Teams struct {
	Map map[string]*Team `json:"teams"`
}
