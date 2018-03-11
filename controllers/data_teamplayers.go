package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/models"
)

//AddPlayerToTeamParams paramaters from player addition to team query
type AddPlayerToTeamParams struct {
	PlayerID string `json:"playerid"`
	TeamID   string `json:"teamid"`
}

//AddPlayerToTeam add a player to team from both IDs
func (c *DataController) AddPlayerToTeam() {
	params := AddPlayerToTeamParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			if !models.PlayerIDListContains(team.Players, params.PlayerID) {
				team.Players = append(team.Players,
					&models.PlayerData{
						ID:      params.PlayerID,
						Comment: "",
					})
				err = SerializeTeamToDB(b, team)
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

//RemovePlayerFromTeamParams paramaters from player removal from game query
type RemovePlayerFromTeamParams struct {
	PlayerID string `json:"playerid"`
	TeamID   string `json:"teamid"`
}

//RemovePlayerFromTeam remove a player from game from both IDs
func (c *DataController) RemovePlayerFromTeam() {
	params := RemovePlayerFromTeamParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			for i, player := range team.Players {
				if player.ID == params.PlayerID {
					team.Players = append(team.Players[:i], team.Players[i+1:]...)
					break
				}
			}
			err = SerializeTeamToDB(b, team)
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

//SubmitPlayerTeamCommentParams paramaters for player comment submission for team query
type SubmitPlayerTeamCommentParams struct {
	PlayerID string `json:"playerid"`
	TeamID   string `json:"teamid"`
	Comment  string `json:"comment"`
}

//SubmitPlayerTeamComment submit player comment for game from both IDs
func (c *DataController) SubmitPlayerTeamComment() {
	params := SubmitPlayerTeamCommentParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			for _, player := range team.Players {
				if player.ID == params.PlayerID {
					player.Comment = params.Comment
					break
				}
			}
			err = SerializeTeamToDB(b, team)
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
