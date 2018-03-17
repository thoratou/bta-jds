package models

//Post forum post data
type ForumPost struct {
	ID       string `json:"id"`
	PlayerID string `json:"playerid"`
	Content  string `json:"content"`
}
