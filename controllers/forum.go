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
		ID:               fmt.Sprintf("%05d", newPostID),
		PlayerID:         playerid,
		Content:          content,
		CreationDate:     time.Now().Format("02/01/2006 à 15:04:05"),
		ModificationDate: "",
		DeletionDate:     "",
	}
	return newPost
}

//ModifyPost modify an existing post with new content
func ModifyPost(team *models.Team, postid string, content string) {
	post := team.Forum[postid]
	post.Content = content
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//DeletePost mark an existing post as deleted
func DeletePost(team *models.Team, postid string) {
	post := team.Forum[postid]
	post.DeletionDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//RestorePost restore a deleted post
func RestorePost(team *models.Team, postid string) {
	post := team.Forum[postid]
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
	post.DeletionDate = ""
}
