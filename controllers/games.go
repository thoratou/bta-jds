package controllers

import (
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/models"
)

//GameController main page for game register
type GameController struct {
	beego.Controller
}

//Get handle get request
func (c *GameController) Get() {
	db := GetDB()
	var data *models.Games
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		dbData, err := DeserializeAllGamesFromDB(b)
		data = dbData
		return err
	})

	//res, _ := SerializeAllGamesToJSON(data)
	beego.Info(data)
	c.Data["json"] = data
	c.ServeJSON()
}
