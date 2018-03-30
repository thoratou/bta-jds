package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

//SubmitGameNewPostParams paramaters for player forum submission for game query
type SubmitGameNewPostParams struct {
	GameID   string `json:"gameid"`
	NewPost  string `json:"newpost"`
	PlayerID string `json:"playerid"`
}

//SubmitGameNewPost submit player new post for game from both IDs
func (c *DataController) SubmitGameNewPost() {
	params := SubmitGameNewPostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			newPost := CreateNewPost(tx.Bucket([]byte("forumids")), params.NewPost, params.PlayerID)
			game.Forum[newPost.ID] = newPost
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

//SubmitGameModifyPostParams paramaters for player forum submission for game query
type SubmitGameModifyPostParams struct {
	GameID       string `json:"gameid"`
	PostID       string `json:"postid"`
	ModifiedPost string `json:"modifiedpost"`
	PlayerID     string `json:"playerid"`
}

//SubmitGameModifyPost modify player post for game from both IDs
func (c *DataController) SubmitGameModifyPost() {
	params := SubmitGameModifyPostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			post, postexists := game.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				ModifyGamePost(game, params.PostID, params.ModifiedPost)
				err = SerializeGameToDB(b, game)
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

//SubmitGameDeletePostParams paramaters for player forum submission for game query
type SubmitGameDeletePostParams struct {
	GameID   string `json:"gameid"`
	PostID   string `json:"postid"`
	PlayerID string `json:"playerid"`
}

//SubmitGameDeletePost delete player post for game from both IDs
func (c *DataController) SubmitGameDeletePost() {
	params := SubmitGameDeletePostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			post, postexists := game.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				DeleteGamePost(game, params.PostID)
				err = SerializeGameToDB(b, game)
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

//RestoreGamePostParams paramaters for player forum submission for game query
type RestoreGamePostParams struct {
	GameID   string `json:"gameid"`
	PostID   string `json:"postid"`
	PlayerID string `json:"playerid"`
}

//RestoreGamePost restore player post for game from both IDs
func (c *DataController) RestoreGamePost() {
	params := RestoreGamePostParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			post, postexists := game.Forum[params.PostID]
			if postexists && post.PlayerID == params.PlayerID {
				RestoreGamePost(game, params.PostID)
				err = SerializeGameToDB(b, game)
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
