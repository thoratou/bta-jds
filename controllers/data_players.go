package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/models"
)

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
