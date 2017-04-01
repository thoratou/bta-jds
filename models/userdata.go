package models

//UserData user data in DB
type UserData struct {
	SHAPassword []byte `json:"shapassword"`
	PlayerID    string `json:"playerid"`
}
