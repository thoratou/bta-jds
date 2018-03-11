package main_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds-jds/controllers"
	"github.com/thoratou/organize-jds-jds/models"
)

func TestCreateUsers(t *testing.T) {
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("users"))
		tx.CreateBucketIfNotExists([]byte("players"))

		createDefaultUser(tx, "toto1@cgi.com")
		createDefaultUser(tx, "toto2@cgi.com")
		createDefaultUser(tx, "toto3@cgi.com")
		return nil
	})
}

func createDefaultUser(tx *bolt.Tx, mail string) {
	userData := models.UserData{}

	users := tx.Bucket([]byte("users"))
	existingData := users.Get([]byte(mail))
	if existingData != nil {
		if err := json.Unmarshal(existingData, &userData); err != nil {
			return
		}
	} else {
		players := tx.Bucket([]byte("players"))
		userData.PlayerID = controllers.CreatePlayer(players, mail)
	}

	userData.SHAPassword = controllers.ConvertToSHA1("default")

	//put user data in all cases to at least reset password
	updatedUser, _ := json.Marshal(userData)
	users.Put([]byte(mail), updatedUser)
}
