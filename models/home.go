package models

//SignIn parameters
type SignIn struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

//SignUp parameters
type SignUp struct {
	User string `json:"user"`
}

//UserData user data in DB
type UserData struct {
	SHAPassword []byte `json:"shapassword"`
}
