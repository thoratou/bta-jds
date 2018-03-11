package controllers

import (
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/models"
)

//Data global data
type Data struct {
	Games           map[string]*models.Game   `json:"games"`
	Teams           map[string]*models.Team   `json:"teams"`
	Players         map[string]*models.Player `json:"players"`
	CurrentPlayerID string                    `json:"currentplayerid"`
}

//DataController main page for game register
type DataController struct {
	beego.Controller
}

//Get handle get request
func (c *DataController) Get() {
	db := GetDB()
	data := &Data{}

	//games
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		dbData, err := DeserializeAllGamesFromDB(b)
		data.Games = dbData.Map
		return err
	})

	//teams
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		dbData, err := DeserializeAllTeamsFromDB(b)
		data.Teams = dbData.Map
		return err
	})

	//players
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("players"))
		dbData, err := DeserializeAllPlayersFromDB(b)
		data.Players = dbData.Map
		return err
	})

	//players
	mail := c.Ctx.Input.Param(":splat")
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		id, err := DeserializeCurrentPlayerIDFromDB(b, mail)
		data.CurrentPlayerID = id
		return err
	})

	//res, _ := SerializeAllGamesToJSON(data)
	beego.Info(data)
	c.Data["json"] = data
	c.ServeJSON()
}
