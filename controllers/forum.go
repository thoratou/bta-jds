package controllers

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/models"
)

//CreateNewPost create a new post
func CreateNewPost(bucket *bolt.Bucket, content string, playerid string) *models.ForumPost {
	newPostID, _ := bucket.NextSequence()
	newPost := &models.ForumPost{
		ID:       fmt.Sprintf("%03d", newPostID),
		PlayerID: playerid,
		Content:  content,
	}
	return newPost
}
