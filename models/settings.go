package models

//Settings company settings
type Settings struct {
	Site           string `json:"site"`
	CSS            string `json:"css"`
	CompanyName    string `json:"companyName"`
	MailExtension  string `json:"mailExtension"`
	SenderMail     string `json:"senderMail"`
	SenderPassword string `json:"senderPassword"`
	Contact        string `json:"contact"`
	Logo           string `json:"logo"`
}
