package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/models"
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

//AddPlayerToGameParams paramaters from player addition to game query
type AddPlayerToGameParams struct {
	PlayerID string `json:"playerid"`
	GameID   string `json:"gameid"`
}

//AddPlayerToGame add a player to game from both IDs
func (c *DataController) AddPlayerToGame() {
	params := AddPlayerToGameParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			if !models.PlayerIDListContains(game.Players, params.PlayerID) {
				game.Players = append(game.Players,
					&models.PlayerData{
						ID:      params.PlayerID,
						Comment: "",
					})
				err = SerializeGameToDB(b, game)
			}
		}
		return err
	})

	if err != nil {
		c.Ctx.Output.SetStatus(501)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.SetStatus(200)

}

//RemovePlayerFromGameParams paramaters from player removal from game query
type RemovePlayerFromGameParams struct {
	PlayerID string `json:"playerid"`
	GameID   string `json:"gameid"`
}

//RemovePlayerFromGame remove a player from game from both IDs
func (c *DataController) RemovePlayerFromGame() {
	params := RemovePlayerFromGameParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			for i, player := range game.Players {
				if player.ID == params.PlayerID {
					game.Players = append(game.Players[:i], game.Players[i+1:]...)
					break
				}
			}
			err = SerializeGameToDB(b, game)
		}
		return err
	})

	if err != nil {
		c.Ctx.Output.SetStatus(501)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.SetStatus(200)

}

//SubmitPlayerGameCommentParams paramaters for player comment submission for game query
type SubmitPlayerGameCommentParams struct {
	PlayerID string `json:"playerid"`
	GameID   string `json:"gameid"`
	Comment  string `json:"comment"`
}

//SubmitPlayerGameComment submit player comment for game from both IDs
func (c *DataController) SubmitPlayerGameComment() {
	params := SubmitPlayerGameCommentParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			for _, player := range game.Players {
				if player.ID == params.PlayerID {
					player.Comment = params.Comment
					break
				}
			}
			err = SerializeGameToDB(b, game)
		}
		return err
	})

	if err != nil {
		c.Ctx.Output.SetStatus(501)
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.SetStatus(200)

}
