package controllers

import (
	"encoding/json"

	"bytes"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/models"
)

//HomeController home pase with authentication
type HomeController struct {
	beego.Controller
}

//Get handle get request
func (c *HomeController) Get() {
	c.Data["Website"] = GetSettings().Site
	c.Data["Email"] = GetSettings().Contact
	c.Data["CSS"] = GetSettings().CSS
	c.Data["CompanyName"] = GetSettings().CompanyName
	c.Data["MailExtension"] = GetSettings().MailExtension
	c.Data["SendMail"] = GetSettings().SenderMail
	c.Data["Contact"] = GetSettings().Contact
	c.Data["Logo"] = GetSettings().Logo
	c.TplName = "index.html"
	c.Render()
}

//SignInQuery handle login request (post)
//
//   req: POST /signin {"user": "your username", "password": "your awesome password"}
//   res: 200 SignInSuccessful
//   res: 403 Invalid username or password
//   res: 404 Missing username
//   res: 405 Missing password
//   res: 500 Unknown error
func (c *HomeController) SignInQuery() {
	data := models.SignIn{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	user := data.User
	if user == "" {
		c.Ctx.Output.SetStatus(404)
		return
	}

	mail := user + "@" + GetSettings().MailExtension

	password := data.Password
	if password == "" {
		c.Ctx.Output.SetStatus(405)
		return
	}

	db := GetDB()
	userData := models.UserData{}

	userExists := false
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v := b.Get([]byte(mail))
		if len(v) > 0 {
			if err := json.Unmarshal(v, &userData); err == nil {
				userExists = true
			} else {
				return err
			}
		}

		return nil
	})

	beego.Info("input: ", password, ", toSHA1: ", ConvertToSHA1(password), ", db: ", userData.SHAPassword)

	if err != nil {
		c.Ctx.Output.SetStatus(501)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	if !userExists || !bytes.Equal(userData.SHAPassword, ConvertToSHA1(password)) {
		c.Ctx.Output.SetStatus(403)
		return
	}

	c.Ctx.Output.SetStatus(200)
}

//SignUpQuery handle user creation and password reset (post)
//
//   req: POST /signup {"user": "your username"}
//   res: 200 SignInSuccessful
//   res: 404 Missing username
//   res: 500 Unknown error
func (c *HomeController) SignUpQuery() {
	data := models.SignUp{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
		c.Ctx.Output.SetStatus(502)
		return
	}

	user := data.User
	if user == "" {
		c.Ctx.Output.SetStatus(404)
		return
	}

	mail := user + "@" + GetSettings().MailExtension

	db := GetDB()
	userData := models.UserData{}
	newPassword := CreateRandomPassword()

	dbErr := db.Update(func(tx *bolt.Tx) error {

		users := tx.Bucket([]byte("users"))
		existingData := users.Get([]byte(mail))
		if existingData != nil {
			if err := json.Unmarshal(existingData, &userData); err != nil {
				return err
			}
		} else {
			players := tx.Bucket([]byte("players"))
			userData.PlayerID = CreatePlayer(players, mail)
		}

		beego.Info("random: ", newPassword, ", toSHA1: ", userData.SHAPassword)
		userData.SHAPassword = ConvertToSHA1(newPassword)

		//put user data in all cases to at least reset password
		updatedUser, err := json.Marshal(userData)
		if err != nil {
			return err
		}
		return users.Put([]byte(mail), updatedUser)
	})

	if dbErr != nil {
		c.Ctx.Output.SetStatus(501)
		c.Ctx.Output.Body([]byte(dbErr.Error()))
		return
	}

	if err := SendMail(mail, newPassword); err != nil {
		c.Ctx.Output.SetStatus(503)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.SetStatus(200)
}
