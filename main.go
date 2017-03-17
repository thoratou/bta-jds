package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/controllers"
)

func main() {
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("users"))
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("games"))
		if b == nil {
			createdBucket, _ := tx.CreateBucket([]byte("games"))
			controllers.CreateTeamGame(createdBucket, "BabyFoot", 2)
			controllers.CreateIndividualGame(createdBucket, "Badminton")
			controllers.CreateTeamGame(createdBucket, "Badminton Double", 2)
			controllers.CreateTeamGame(createdBucket, "Basketball", 5)
			controllers.CreateTeamGame(createdBucket, "Beach Volley", 3)
			controllers.CreateTeamGame(createdBucket, "Boules Carrées", 2)
			controllers.CreateIndividualGame(createdBucket, "Bowling")
			controllers.CreateIndividualGame(createdBucket, "Color Race")
			controllers.CreateTeamGame(createdBucket, "Course Orientation", 2)
			controllers.CreateIndividualGame(createdBucket, "Cross")
			controllers.CreateIndividualGame(createdBucket, "Danse Zumba")
			controllers.CreateTeamGame(createdBucket, "Fléchettes", 2)
			controllers.CreateTeamGame(createdBucket, "Football à 5", 5)
			controllers.CreateTeamGame(createdBucket, "Football Féminin", 7)
			controllers.CreateTeamGame(createdBucket, "Football Masculin", 7)
			controllers.CreateIndividualGame(createdBucket, "Geocaching")
			controllers.CreateTeamGame(createdBucket, "Golf", 2)
			controllers.CreateTeamGame(createdBucket, "Handball", 6)
			controllers.CreateIndividualGame(createdBucket, "Jeux de l'Esprit")
			controllers.CreateIndividualGame(createdBucket, "Karting")
			controllers.CreateTeamGame(createdBucket, "Kayak", 2)
			controllers.CreateTeamGame(createdBucket, "Laser Quest", 3)
			controllers.CreateIndividualGame(createdBucket, "Marche")
			controllers.CreateIndividualGame(createdBucket, "Nautathlon")
			controllers.CreateTeamGame(createdBucket, "Padel", 2)
			controllers.CreateTeamGame(createdBucket, "Pétanque", 3)
			controllers.CreateIndividualGame(createdBucket, "Photomathon")
			controllers.CreateIndividualGame(createdBucket, "Poker")
			controllers.CreateTeamGame(createdBucket, "Rugby a 7", 7)
			controllers.CreateIndividualGame(createdBucket, "Soirée de cloture")
			controllers.CreateIndividualGame(createdBucket, "Sports Co Decouverte")
			controllers.CreateTeamGame(createdBucket, "Tennis", 2)
			controllers.CreateIndividualGame(createdBucket, "Tennis de Table")
			controllers.CreateIndividualGame(createdBucket, "Tir à l'Arc")
			controllers.CreateIndividualGame(createdBucket, "Tir Au Pistolet")
			controllers.CreateTeamGame(createdBucket, "Ultimate Frisbee", 5)
			controllers.CreateTeamGame(createdBucket, "Voile", 2)
			controllers.CreateTeamGame(createdBucket, "VolleyBall", 6)
			controllers.CreateIndividualGame(createdBucket, "VTT 15 km")
			controllers.CreateTeamGame(createdBucket, "VTT Nocturne Relais", 2)
		}
		return nil
	})

	controllers.SetDB(db)

	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/signin", &controllers.HomeController{}, "post:SignInQuery")
	beego.Router("/signup", &controllers.HomeController{}, "post:SignUpQuery")

	beego.Router("/games/*", &controllers.GameController{})

	beego.Run()
}
