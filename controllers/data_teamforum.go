package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

//SubmitTeamNewPostParams paramaters for player forum submission for team query
type SubmitTeamNewPostParams struct {
	TeamID   string `json:"teamid"`
	NewPost  string `json:"newpost"`
	PlayerID string `json:"playerid"`
}

//SubmitTeamNewPost submit player new post for team from both IDs
func (c *DataController) SubmitTeamNewPost() {
	params := SubmitTeamNewPostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			newPost := CreateNewPost(tx.Bucket([]byte("forumids")), params.NewPost, params.PlayerID)
			team.Forum[newPost.ID] = newPost
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
