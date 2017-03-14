package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/astaxie/beego"
)

//HomeController home pase with authentication
type HomeController struct {
	beego.Controller
}

//Get handle get request
func (c *HomeController) Get() {
	c.Data["Website"] = "cgi-jds.com"
	c.Data["Email"] = "thoratou@gmail.com"
	c.TplName = "index.html"
	c.Render()
}

//SignInQuery handle login request (post)
//
//   req: POST /signin/user {"Password": "your awesome password"}
//   res: 200 SignInSuccessful
//   res: 403 Invalid username or password
//   res: 404 Missing username (client only error)
//   res: 405 Missing password
//   res: Unknown error
func (c *HomeController) SignInQuery() {
	data := struct {
		User     string
		Password string
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	user := data.User
	if user == "" {
		c.Ctx.Output.SetStatus(403)
		return
	}

	password := data.Password
	if password == "" {
		c.Ctx.Output.SetStatus(405)
		return
	}

	//user+password checks todo

	c.Ctx.Output.SetStatus(200)
}

//SignInReply handle login reply (get)
//
//   req: GET /signin/user
//   res: 200 {"User": "user", "HasError": true, "Error": "Whatever"}
/*
func (c *HomeController) SignInReply() {
		res := struct {
			User     string `json:"user"`
			HasError bool   `json:"haserror"`
			Error    string `json:"error"`
		}{
			"dummyuser",
			true,
			"Error",
		}
		c.Data["json"] = &res
		c.ServeJSON()
}
*/

//SignUpQuery handle password reset or account creation
func (c *HomeController) SignUpQuery() {
	data := struct {
		User string
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	user := data.User
	if user == "" {
		c.Ctx.Output.SetStatus(403)
		return
	}

	//DB check to do

	c.sendMail(user + "@cgi.com")

	c.Ctx.Output.SetStatus(200)
}

func (c *HomeController) sendMail(toAddress string) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"cgijds2017@gmail.com",
		"04mai2013",
		"smtp.gmail.com",
	)

	from := mail.Address{
		Name:    "CGI Jeux de Sophia",
		Address: "cgijds2017@gmail.com",
	}
	to := mail.Address{
		Name:    "",
		Address: toAddress,
	}
	title := "Inscription site CGI Jeux de Sophia"

	body := "Merci de votre inscription sur le site CGI des jeux de Sophia\r\nVotre nouveau mot de passe est ..."

	// Fill header data
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"cgijds2017@gmail.com",
		[]string{toAddress},
		[]byte(message),
	)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}
}
