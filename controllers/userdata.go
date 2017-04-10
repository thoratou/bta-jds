package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/bta-jds/models"
)

//DeserializeCurrentPlayerIDFromDB deserialize current player ID from mail
func DeserializeCurrentPlayerIDFromDB(bucket *bolt.Bucket, mail string) (string, error) {
	v := bucket.Get([]byte(mail))
	userData := &models.UserData{}
	err := json.Unmarshal(v, userData)
	beego.Info("mail:", mail)
	beego.Info("userData:", userData)
	beego.Info("playerID:", userData.PlayerID)
	return userData.PlayerID, err
}
