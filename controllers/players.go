package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/thoratou/bta-jds/models"
)

//CreatePlayer create a player
func CreatePlayer(bucket *bolt.Bucket, mail string) string {
	newPlayer := &models.Player{
		FirstName: "",
		LastName:  "",
		Mail:      mail,
	}

	newPlayerID, _ := bucket.NextSequence()
	idStr, _ := SerializeNewPlayerToDB(bucket, newPlayerID, newPlayer)
	return idStr
}

//SerializeNewPlayerToDB serialize player data to database
func SerializeNewPlayerToDB(bucket *bolt.Bucket, id uint64, newPlayer *models.Player) (string, error) {
	idStr := fmt.Sprintf("%03d", id)
	newPlayer.ID = idStr
	if v, err := json.Marshal(newPlayer); err == nil {
		bucket.Put([]byte(idStr), v)
	} else {
		return "", err
	}
	return idStr, nil
}

//DeserializeAllPlayersFromDB deserialize all players data from database
func DeserializeAllPlayersFromDB(bucket *bolt.Bucket) (*models.Players, error) {
	data := &models.Players{
		Map: make(map[string]*models.Player),
	}
	err := bucket.ForEach(func(k, v []byte) error {

		idStr := string(k)

		player := &models.Player{}
		err := json.Unmarshal(v, player)
		if err == nil {
			data.Map[idStr] = player
		}
		return err
	})

	return data, err
}
