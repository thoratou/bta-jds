package controllers

import (
	"encoding/json"

	"fmt"

	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/models"
)

//CreateTeamGame create a team game with a minimum number of players
func CreateTeamGame(bucket *bolt.Bucket, name string, minPlayers int, description []string) {
	newGameID, _ := bucket.NextSequence()
	newGame := &models.Game{
		Name:        name,
		TeamGame:    true,
		MinPlayers:  minPlayers,
		Description: description,
		Players:     []models.PlayerData{},
	}

	SerializeNewGameToDB(bucket, newGameID, newGame)
}

//CreateIndividualGame create a individual game
func CreateIndividualGame(bucket *bolt.Bucket, name string, description []string) {
	newGame := &models.Game{
		Name:        name,
		TeamGame:    false,
		MinPlayers:  0,
		Description: description,
		Players:     []models.PlayerData{},
	}

	newGameID, _ := bucket.NextSequence()
	SerializeNewGameToDB(bucket, newGameID, newGame)
}

//SerializeNewGameToDB serialize game data to database
func SerializeNewGameToDB(bucket *bolt.Bucket, id uint64, newGame *models.Game) error {
	idStr := fmt.Sprintf("%03d", id)
	newGame.ID = idStr
	if v, err := json.Marshal(newGame); err == nil {
		bucket.Put([]byte(idStr), v)
	} else {
		return err
	}
	return nil
}

//SerializeGameToDB serialize game data to database
func SerializeGameToDB(bucket *bolt.Bucket, game *models.Game) error {
	if v, err := json.Marshal(game); err == nil {
		bucket.Put([]byte(game.ID), v)
	} else {
		return err
	}
	return nil
}

//DeserializeAllGamesFromDB deserialize all games data from database
func DeserializeAllGamesFromDB(bucket *bolt.Bucket) (*models.Games, error) {
	data := &models.Games{
		Map: make(map[string]*models.Game),
	}
	err := bucket.ForEach(func(k, v []byte) error {

		idStr := string(k)

		game := &models.Game{}
		err := json.Unmarshal(v, game)
		if err == nil {
			data.Map[idStr] = game
		}
		return err
	})

	return data, err
}

//DeserializeGameFromDB deserialize game data  with game ID from database
func DeserializeGameFromDB(bucket *bolt.Bucket, gameID string) (*models.Game, error) {
	data := &models.Game{}
	v := bucket.Get([]byte(gameID))

	err := json.Unmarshal(v, data)
	return data, err
}
