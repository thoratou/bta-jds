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

//ModifyTeamPost modify an existing post with new content
func ModifyTeamPost(team *models.Team, postid string, content string) {
	post := team.Forum[postid]
	post.Content = content
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//DeleteTeamPost mark an existing post as deleted
func DeleteTeamPost(team *models.Team, postid string) {
	post := team.Forum[postid]
	post.DeletionDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//RestoreTeamPost restore a deleted post
func RestoreTeamPost(team *models.Team, postid string) {
	post := team.Forum[postid]
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
	post.DeletionDate = ""
}

//ModifyGamePost modify an existing post with new content
func ModifyGamePost(game *models.Game, postid string, content string) {
	post := game.Forum[postid]
	post.Content = content
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//DeleteGamePost mark an existing post as deleted
func DeleteGamePost(game *models.Game, postid string) {
	post := game.Forum[postid]
	post.DeletionDate = time.Now().Format("02/01/2006 à 15:04:05")
}

//RestoreGamePost restore a deleted post
func RestoreGamePost(game *models.Game, postid string) {
	post := game.Forum[postid]
	post.ModificationDate = time.Now().Format("02/01/2006 à 15:04:05")
	post.DeletionDate = ""
}
