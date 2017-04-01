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
		tx.CreateBucketIfNotExists([]byte("teams"))
		tx.CreateBucketIfNotExists([]byte("players"))

		b := tx.Bucket([]byte("games"))
		if b == nil {
			createdBucket, _ := tx.CreateBucket([]byte("games"))
			controllers.CreateTeamGame(createdBucket, "BabyFoot", 2, []string{
				"Matchs de qualification en 2 contre 2",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Mardi 14 juin 2016",
				"Dates des qualifications : Lundi 13, Mardi 14, Mercredi 15, Jeudi 16 Juin 2016 de 12h15 à 13h45",
				"Date de la finale : Vendredi 17 Juin 2016 de 12h15 à 13h45",
			})
			controllers.CreateIndividualGame(createdBucket, "Badminton", []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Jeudi 16 Juin 2016",
				"Date de l'épreuve : Samedi 18 Juin 2016 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Badminton Double", 2, []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Jeudi 16 Juin 2016",
				"Date de l'épreuve : Samedi 18 Juin 2016 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Basketball", 5, []string{
				"Épreuve mixte mais aucune obligation de présence d’une fille dans le cinq entrant sur le terrain",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Mercredi 1 Juin  2016",
				"Dates des qualifications : Lundi 30 Mai, Mardi 31 Mai,  Mercredi 01, Jeudi 02, Vendredi 3, Mardi 7 juin de 17h30 à 21h30",
				"Date des phases finales : Mercredi 8 juin de 18h à 21h30",
			})
			controllers.CreateTeamGame(createdBucket, "Beach Volley", 3, []string{
				"Tournoi mixte sur sable 3 x 3",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Lundi 06 Juin 2016",
				"Dates des qualifications : Mardi 31, mercredi 01, jeudi 02, lundi 06, mardi 07 et mercredi 08 juinde 18h à 21h00",
				"Date de la finale : Jeudi 09 juin de 18h à 21h30",
			})
			controllers.CreateTeamGame(createdBucket, "Boules Carrées", 2, []string{
				"Épreuve disputée en doublettes formées (2 joueurs, pas forcément d'une même entreprise/école) avec des boules carrées",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Jeudi 16 Juin 2016",
				"Dates des épreuves : Samedi 18 Juin de 8h00 à 20h00",
			})
			controllers.CreateIndividualGame(createdBucket, "Bowling", []string{})
			controllers.CreateIndividualGame(createdBucket, "Color Race", []string{})
			controllers.CreateTeamGame(createdBucket, "Course Orientation", 2, []string{})
			controllers.CreateIndividualGame(createdBucket, "Cross", []string{})
			controllers.CreateIndividualGame(createdBucket, "Danse Zumba", []string{})
			controllers.CreateTeamGame(createdBucket, "Fléchettes", 2, []string{})
			controllers.CreateTeamGame(createdBucket, "Football à 5", 5, []string{})
			controllers.CreateTeamGame(createdBucket, "Football Féminin", 7, []string{})
			controllers.CreateTeamGame(createdBucket, "Football Masculin", 7, []string{})
			controllers.CreateIndividualGame(createdBucket, "Geocaching", []string{})
			controllers.CreateTeamGame(createdBucket, "Golf", 2, []string{})
			controllers.CreateTeamGame(createdBucket, "Handball", 6, []string{})
			controllers.CreateIndividualGame(createdBucket, "Jeux de l'Esprit", []string{})
			controllers.CreateIndividualGame(createdBucket, "Karting", []string{})
			controllers.CreateTeamGame(createdBucket, "Kayak", 2, []string{})
			controllers.CreateTeamGame(createdBucket, "Laser Quest", 3, []string{})
			controllers.CreateIndividualGame(createdBucket, "Marche", []string{})
			controllers.CreateIndividualGame(createdBucket, "Nautathlon", []string{})
			controllers.CreateTeamGame(createdBucket, "Padel", 2, []string{})
			controllers.CreateTeamGame(createdBucket, "Pétanque", 3, []string{})
			controllers.CreateIndividualGame(createdBucket, "Photomathon", []string{})
			controllers.CreateIndividualGame(createdBucket, "Poker", []string{})
			controllers.CreateTeamGame(createdBucket, "Rugby a 7", 7, []string{})
			controllers.CreateIndividualGame(createdBucket, "Soirée de cloture", []string{})
			controllers.CreateIndividualGame(createdBucket, "Sports Co Decouverte", []string{})
			controllers.CreateTeamGame(createdBucket, "Tennis", 2, []string{})
			controllers.CreateIndividualGame(createdBucket, "Tennis de Table", []string{})
			controllers.CreateIndividualGame(createdBucket, "Tir à l'Arc", []string{})
			controllers.CreateIndividualGame(createdBucket, "Tir Au Pistolet", []string{})
			controllers.CreateTeamGame(createdBucket, "Ultimate Frisbee", 5, []string{})
			controllers.CreateTeamGame(createdBucket, "Voile", 2, []string{})
			controllers.CreateTeamGame(createdBucket, "VolleyBall", 6, []string{})
			controllers.CreateIndividualGame(createdBucket, "VTT 15 km", []string{})
			controllers.CreateTeamGame(createdBucket, "VTT Nocturne Relais", 2, []string{})
		}
		return nil
	})

	controllers.SetDB(db)

	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/signin", &controllers.HomeController{}, "post:SignInQuery")
	beego.Router("/signup", &controllers.HomeController{}, "post:SignUpQuery")

	beego.Router("/games/*", &controllers.DataController{}, "get:Get")
	beego.Router("/addPlayerToGame", &controllers.DataController{}, "post:AddPlayerToGame")
	beego.Router("/removePlayerFromGame", &controllers.DataController{}, "post:RemovePlayerFromGame")
	beego.Router("/submitPlayerGameComment", &controllers.DataController{}, "post:SubmitPlayerGameComment")

	beego.Run()
}
