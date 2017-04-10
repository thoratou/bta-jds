package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

//SubmitPlayerDataParams paramaters for player data submission
type SubmitPlayerDataParams struct {
	PlayerID  string `json:"playerid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//SubmitPlayerData submit player data
func (c *DataController) SubmitPlayerData() {
	params := SubmitPlayerDataParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("players"))
		player, err := DeserializePlayerFromDB(b, params.PlayerID)
		if err == nil {
			player.FirstName = params.FirstName
			player.LastName = params.LastName
			err = SerializePlayerToDB(b, player)
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
