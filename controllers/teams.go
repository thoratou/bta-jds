package controllers

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/models"
)

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
