package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

//AddTeamToGameParams paramaters from team addition to game query
type AddTeamToGameParams struct {
	TeamName  string `json:"teamname"`
	ManagerID string `json:"managerid"`
	GameID    string `json:"gameid"`
}

//AddTeamToGame add a team to game from both IDs
func (c *DataController) AddTeamToGame() {
	params := AddTeamToGameParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		teamBucket := tx.Bucket([]byte("teams"))
		teamID := CreateTeam(teamBucket, params.TeamName, params.ManagerID, params.GameID)

		b := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(b, params.GameID)
		if err == nil {
			game.Teams = append(game.Teams, teamID)
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

//RemoveTeamFromGameParams paramaters from team removal from game query
type RemoveTeamFromGameParams struct {
	TeamID    string `json:"teamid"`
	ManagerID string `json:"managerid"`
	GameID    string `json:"gameid"`
}

//RemoveTeamFromGame remove a team from game from both IDs
func (c *DataController) RemoveTeamFromGame() {
	params := RemoveTeamFromGameParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		//game part
		gameBucket := tx.Bucket([]byte("games"))
		game, err := DeserializeGameFromDB(gameBucket, params.GameID)
		if err == nil {
			for i, teamid := range game.Teams {
				if teamid == params.TeamID {
					game.Teams = append(game.Teams[:i], game.Teams[i+1:]...)
					break
				}
			}
			err = SerializeGameToDB(gameBucket, game)
		}

		if err == nil {
			//team part
			teamBucket := tx.Bucket([]byte("teams"))
			team, err := DeserializeTeamFromDB(teamBucket, params.TeamID)
			if err == nil {
				if len(team.Players) > 0 {
					//only flagged as removed to keep backup data
					team.Removed = true
					err = SerializeTeamToDB(teamBucket, team)
				} else {
					//surely created by mistake, removed it
					teamBucket.Delete([]byte(team.ID))
				}
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

//ChangeTeamNameParams paramaters from team name change query
type ChangeTeamNameParams struct {
	TeamID    string `json:"teamid"`
	TeamName  string `json:"teamname"`
	ManagerID string `json:"managerid"`
}

//ChangeTeamName remove a team from game from both IDs
func (c *DataController) ChangeTeamName() {
	params := ChangeTeamNameParams{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}

	db := GetDB()
	err := db.Update(func(tx *bolt.Tx) error {
		teamBucket := tx.Bucket([]byte("teams"))
		team, err := DeserializeTeamFromDB(teamBucket, params.TeamID)
		if err == nil {
			team.Name = params.TeamName
			err = SerializeTeamToDB(teamBucket, team)
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
