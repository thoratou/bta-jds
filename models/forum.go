package models

//ForumPost forum post data
type ForumPost struct {
	ID               string `json:"id"`
	PlayerID         string `json:"playerid"`
	Content          string `json:"content"`
	CreationDate     string `json:"creationdate"`
	ModificationDate string `json:"modificationdate"`
	DeletionDate     string `json:"deletiondate"`
}
