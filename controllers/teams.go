package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/thoratou/bta-jds/models"
)

//CreateTeam create a team
func CreateTeam(bucket *bolt.Bucket, name string, managerID string, gameID string) string {
	newTeam := &models.Team{
		Name:      name,
		ManagerID: managerID,
		Players: []*models.PlayerData{ //add manager by default
			&models.PlayerData{
				ID:      managerID,
				Comment: "",
			},
		},
		GameID:  gameID,
		Comment: "",
		Removed: false,
	}

	newTeamID, _ := bucket.NextSequence()
	idStr, _ := SerializeNewTeamToDB(bucket, newTeamID, newTeam)
	return idStr
}

//SerializeNewTeamToDB serialize team data to database
func SerializeNewTeamToDB(bucket *bolt.Bucket, id uint64, newTeam *models.Team) (string, error) {
	idStr := fmt.Sprintf("%03d", id)
	newTeam.ID = idStr
	if v, err := json.Marshal(newTeam); err == nil {
		bucket.Put([]byte(idStr), v)
	} else {
		return "", err
	}
	return idStr, nil
}

//SerializeTeamToDB serialize team data to database
func SerializeTeamToDB(bucket *bolt.Bucket, team *models.Team) error {
	if v, err := json.Marshal(team); err == nil {
		bucket.Put([]byte(team.ID), v)
	} else {
		return err
	}
	return nil
}

//DeserializeAllTeamsFromDB deserialize all teams data from database
func DeserializeAllTeamsFromDB(bucket *bolt.Bucket) (*models.Teams, error) {
	data := &models.Teams{
		Map: make(map[string]*models.Team),
	}
	err := bucket.ForEach(func(k, v []byte) error {

		idStr := string(k)

		team := &models.Team{}
		err := json.Unmarshal(v, team)
		if err == nil {
			data.Map[idStr] = team
		}
		return err
	})

	return data, err
}

//DeserializeTeamFromDB deserialize team data  with game ID from database
func DeserializeTeamFromDB(bucket *bolt.Bucket, teamID string) (*models.Team, error) {
	data := &models.Team{}
	v := bucket.Get([]byte(teamID))

	err := json.Unmarshal(v, data)
	return data, err
}
