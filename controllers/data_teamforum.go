package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

//PostError custom error handler for posts
type PostError struct {
	what string
}

//NewPostError initializer for post error
func NewPostError(what string) *PostError {
	err := &PostError{
		what: what,
	}
	return err
}

func (e *PostError) Error() string {
	return e.what
}

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

//SubmitTeamModifyPostParams paramaters for player forum submission for team query
type SubmitTeamModifyPostParams struct {
	TeamID       string `json:"teamid"`
	PostID       string `json:"postid"`
	ModifiedPost string `json:"modifiedpost"`
	PlayerID     string `json:"playerid"`
}

//SubmitTeamModifyPost modify player post for team from both IDs
func (c *DataController) SubmitTeamModifyPost() {
	params := SubmitTeamModifyPostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			post, postexists := team.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				ModifyTeamPost(team, params.PostID, params.ModifiedPost)
				err = SerializeTeamToDB(b, team)
			} else {
				return NewPostError("invalid data parameters")
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

//SubmitTeamDeletePostParams paramaters for player forum submission for team query
type SubmitTeamDeletePostParams struct {
	TeamID   string `json:"teamid"`
	PostID   string `json:"postid"`
	PlayerID string `json:"playerid"`
}

//SubmitTeamDeletePost delete player post for team from both IDs
func (c *DataController) SubmitTeamDeletePost() {
	params := SubmitTeamDeletePostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			post, postexists := team.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				DeleteTeamPost(team, params.PostID)
				err = SerializeTeamToDB(b, team)
			} else {
				return NewPostError("invalid data parameters")
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

//RestoreTeamPostParams paramaters for player forum submission for team query
type RestoreTeamPostParams struct {
	TeamID   string `json:"teamid"`
	PostID   string `json:"postid"`
	PlayerID string `json:"playerid"`
}

//RestoreTeamPost restore player post for team from both IDs
func (c *DataController) RestoreTeamPost() {
	params := RestoreTeamPostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(b, params.TeamID)
		if err == nil {
			post, postexists := team.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				RestoreTeamPost(team, params.PostID)
				err = SerializeTeamToDB(b, team)
			} else {
				return NewPostError("invalid data parameters")
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
