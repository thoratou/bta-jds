package controllers

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/models"
)

//CreateNewPost create a new post
func CreateNewPost(bucket *bolt.Bucket, content string, playerid string) *models.ForumPost {
	newPostID, _ := bucket.NextSequence()
	newPost := &models.ForumPost{
		ID:               fmt.Sprintf("%03d", newPostID),
		PlayerID:         playerid,
		Content:          content,
		CreationDate:     time.Now().Format("02/01/2006 à 15:04:05"),
		ModificationDate: "",
	}
	return newPost
}
